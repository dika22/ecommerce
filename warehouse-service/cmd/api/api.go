package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warehouse-service/cmd/middleware"
	"warehouse-service/internal/domain/stock/delivery"
	"warehouse-service/internal/domain/stock/usecase"
	"warehouse-service/metrics"
	"warehouse-service/package/logger"

	dw "warehouse-service/internal/domain/warehouse/delivery"
	uw "warehouse-service/internal/domain/warehouse/usecase"

	"os/signal"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
)

const CmdServeHTTP = "serve-http"

type HTTP struct{
	us usecase.IStock
	uw uw.IWarehouse
}

func (h HTTP) ServeAPI(c *cli.Context) error  {
	if err := logger.SetLogger(); err != nil {
		log.Printf("error logger %v", err)
	}
	// Register metrics
	metrics.Register()

	e := echo.New();
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	warehouseAPI := e.Group("api/v1")
	warehouseAPI.Use(middleware.LoggerMiddleware)
	warehouseAPI.Use(middleware.MonitoringMiddleware)

	delivery.NewStockHTTP(warehouseAPI, h.us)
	dw.NewWarehouseHTTP(warehouseAPI, h.uw)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%v", 3004)); err != nil {
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

func ServeAPI(us usecase.IStock, whu uw.IWarehouse) []*cli.Command {
	h := &HTTP{us: us, uw: whu}
	return []*cli.Command{
		{
			Name: CmdServeHTTP,
			Usage: "Serve Document Service",
			Action: h.ServeAPI,
		},
	}
}