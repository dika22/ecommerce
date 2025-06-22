package delivery

import (
	"net/http"
	"warehouse-service/internal/domain/stock/usecase"
	"warehouse-service/package/response"
	"warehouse-service/package/structs"

	"github.com/labstack/echo/v4"
)

type StockHTTP struct{
	uc usecase.IStock
}
func (h StockHTTP) AddStock(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestAddStock{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.AddStock(ctx, req); err != nil {
		return  response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, req, "success add stock")
}

func (h StockHTTP) TransferStock(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestTransferStockProduct{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.TransferStockProduct(ctx, &req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, req, "success transfer stock")
	
}

func (h StockHTTP) GetAvailableStock(c echo.Context) error {
	ctx := c.Request().Context()
	productID := c.Param("product_id")
	resp, err := h.uc.GetAvailableStock(ctx, productID)
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONResponse(c, http.StatusOK, "success", "success get available stock", resp)
	
}

func (h StockHTTP) ReserveStock(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestReserveStock{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.ReserveStock(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, req, "success reserve stock")
}

func (h StockHTTP) ReleaseStock(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestReleaseStock{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.ReleaseStock(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, req, "success release stock")
}


func NewStockHTTP(r *echo.Group, uc usecase.IStock)  {
	u := StockHTTP{uc: uc}

	stockRoute := r.Group("/stocks")
	stockRoute.POST("", u.AddStock).Name = "add.stock"
	stockRoute.POST("/transfer", u.TransferStock).Name = "transfer.stock"
	stockRoute.GET("/available/:product_id", u.GetAvailableStock).Name = "get.available.stock"
	stockRoute.POST("/reserve", u.ReserveStock).Name = "reserve.stock"
	stockRoute.POST("/release", u.ReleaseStock).Name = "release.stock"
}