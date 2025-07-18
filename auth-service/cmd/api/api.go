package cmd

import (
	"auth-service/cmd/middleware"
	sellerDelivery "auth-service/internal/domain/seller/delivery"
	sellerUsecase "auth-service/internal/domain/seller/usecase"
	"auth-service/internal/domain/user/delivery"
	"auth-service/internal/domain/user/usecase"
	"auth-service/metrics"
	"auth-service/package/config"
	"auth-service/package/logger"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct {
	usecase usecase.IUser
	sellerUsecase sellerUsecase.ISeller
	cfg *config.Config
}

func (h HTTP) ServeAPI(c *cli.Context) error {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()

	e := echo.New()
	e.Use(middleware.RatelimitMidleware(h.cfg))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	// Prometheus metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	userAPI := e.Group("api/v1/users")

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		// AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	

	userAPI.Use(middleware.LoggerMiddleware)
	userAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewUserHTTP(userAPI, h.usecase)
	sellerDelivery.NewSellerHTTP(userAPI, h.sellerUsecase)

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

func ServeAPI(usecase usecase.IUser, sellerUsecase sellerUsecase.ISeller, cfg *config.Config) []*cli.Command {
	h := &HTTP{usecase: usecase, sellerUsecase: sellerUsecase, cfg: cfg}
	return []*cli.Command{
		{
			Name:   CmdServeHTTP,
			Usage:  "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}
