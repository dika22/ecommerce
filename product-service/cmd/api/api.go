package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"product-service/cmd/middleware"
	"product-service/internal/domain/product/delivery"
	"product-service/internal/domain/product/usecase"
	"time"

	"os/signal"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	usecase usecase.IProduct
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	e := echo.New();

	// Prometheus metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	// Configurable rate limiter
	rl := middleware.NewRateLimiter(5, 1*time.Second, 0.2) // 5 req / sec with 20% jitter
	e.Use(middleware.RateLimiterMiddleware(rl))

	productAPI := e.Group("api/v1/products")
	productAPI.Use(echoMiddleware.Logger())

	delivery.NewProductHTTP(productAPI, h.usecase)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3001)); err != nil {
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

func ServeAPI(usecase usecase.IProduct) []*cli.Command {
	h := &HTTP{usecase: usecase}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}