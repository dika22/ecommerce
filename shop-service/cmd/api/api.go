package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"shop-service/cmd/middleware"
	"shop-service/internal/domain/shop/delivery"
	"shop-service/internal/domain/shop/usecase"
	"shop-service/package/logger"
	"time"

	"os/signal"

	"shop-service/metrics"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	usecase usecase.IShop
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()
	
	e := echo.New();
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	shopAPI := e.Group("api/v1")
	shopAPI.Use(middleware.LoggerMiddleware)
	shopAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewShopHTTP(shopAPI, h.usecase)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3003)); err != nil {
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

func ServeAPI(usecase usecase.IShop) []*cli.Command {
	h := &HTTP{usecase: usecase}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}