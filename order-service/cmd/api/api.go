package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"order-service/cmd/middleware"
	"order-service/internal/domain/order/delivery"
	"order-service/internal/domain/order/usecase"
	"order-service/metrics"
	"order-service/package/config"
	"order-service/package/logger"
	"os"
	"time"

	"os/signal"

	http_client "order-service/package/http_client"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	usecase usecase.IOrder
	http_clients http_client.HTTPClients
	cacheConf *config.Cache
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()
	
	e := echo.New();

	mon := asynqmon.New(asynqmon.Options{
		RootPath: "/monitoring/tasks",
		RedisConnOpt: asynq.RedisClientOpt{
			Addr: fmt.Sprintf("%s:%s", h.cacheConf.WorkerRedisHost, h.cacheConf.WorkerRedisPort),
		},
	})
	e.Any("/monitoring/tasks/*", echo.WrapHandler(mon))
	
	// Prometheus metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)


	orderAPI := e.Group("api/v1/orders")
	orderAPI.Use(middleware.LoggerMiddleware)
	orderAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewOrderHTTP(orderAPI, h.usecase)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3002)); err != nil {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	return nil	
}

func ServeAPI(usecase usecase.IOrder, http_clients http_client.HTTPClients, cacheConf *config.Cache) []*cli.Command {
	h := &HTTP{usecase: usecase, http_clients: http_clients, cacheConf: cacheConf}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}