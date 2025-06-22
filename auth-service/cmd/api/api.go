package cmd

import (
	"auth-service/cmd/middleware"
	"auth-service/internal/domain/user/delivery"
	"auth-service/internal/domain/user/usecase"
	"auth-service/metrics"
	"auth-service/package/logger"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct {
	usecase usecase.IUser
}

func (h HTTP) ServeAPI(c *cli.Context) error {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	// Prometheus metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	userAPI := e.Group("users/api/v1")
	userAPI.Use(middleware.LoggerMiddleware)
	userAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewUserHTTP(userAPI, h.usecase)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3000)); err != nil {
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

func ServeAPI(usecase usecase.IUser) []*cli.Command {
	h := &HTTP{usecase: usecase}
	return []*cli.Command{
		{
			Name:   CmdServeHTTP,
			Usage:  "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}
