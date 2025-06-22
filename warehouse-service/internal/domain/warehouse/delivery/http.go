package delivery

import (
	"net/http"
	"warehouse-service/internal/domain/warehouse/usecase"
	"warehouse-service/package/response"
	"warehouse-service/package/structs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type WarehouseHTTP struct{
	uc usecase.IWarehouse
}
func (h WarehouseHTTP) AddWarehouse(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestAddWarehouse{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.AddWarehouse(ctx, req); err != nil {
		return err
	}
	return response.JSONSuccess(c, req, "success create warehouse")
}

func (h WarehouseHTTP) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.uc.GetAll(ctx); 
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success get all warehouse")
	
}

func (h WarehouseHTTP) SetWarehouseStatus(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id_warehouse")
	if err := h.uc.SetWarehouseStatus(ctx, cast.ToInt64(id)); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, id, "success set warehouse status")
}

func NewWarehouseHTTP(r *echo.Group, uc usecase.IWarehouse)  {
	u := WarehouseHTTP{uc: uc}

	warehouseRoute := r.Group("/warehouses")
	warehouseRoute.POST("", u.AddWarehouse).Name = "add.warehouse"
	warehouseRoute.PATCH("/:id/status", u.SetWarehouseStatus).Name = "set.warehouse.status"
	warehouseRoute.GET("", u.GetAll).Name = "get.all.warehouse"
}