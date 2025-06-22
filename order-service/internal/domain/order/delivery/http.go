package delivery

import (
	"net/http"
	"order-service/internal/domain/order/usecase"
	"order-service/package/response"
	"order-service/package/structs"

	"github.com/labstack/echo/v4"
)

type orderHTTP struct{
	uc usecase.IOrder
}

// GetAll godoc
//
// @Summary      Request OTP
// @Description  Request OTP for login
// @Tags         auth
// @Status      200  {object}  structs.Response
// @Router       / [get]

func (h orderHTTP) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.uc.GetAll(ctx); 
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success get all order")
}

func (h orderHTTP) GetByID(c echo.Context) error {
	return nil
}

// Create godoc
//
// @Summary      Request OTP
// @Description  Request OTP for login
// @Tags         auth
// @Param        payload  body  structs.RequestCreateOrder  true  "payload"
// @Status      200  {object}  structs.Response
// @Router       / [post]

func (h orderHTTP) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestCreateOrder{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.CreateOrder(ctx, req); err != nil {return err}
	return response.JSONSuccess(c, req, "success create order")
	
}


func NewOrderHTTP(r *echo.Group, uc usecase.IOrder)  {
	u := orderHTTP{uc: uc}
	r.GET("", u.GetAll).Name = "order.get-all"
	r.POST("", u.Create).Name = "order.create"
}