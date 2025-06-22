package delivery

import (
	"net/http"
	"product-service/internal/domain/product/usecase"
	"product-service/package/response"

	"github.com/labstack/echo/v4"
)

type ProductHTTP struct{
	uc usecase.IProduct
}
func (h ProductHTTP) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.uc.GetAll(ctx); 
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success get all product")
}


func NewProductHTTP(r *echo.Group, uc usecase.IProduct)  {
	u := ProductHTTP{uc: uc}
	r.GET("", u.GetAll).Name = "prooduct.get-all"
}