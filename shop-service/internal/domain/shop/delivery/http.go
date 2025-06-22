package delivery

import (
	"net/http"
	"shop-service/internal/domain/shop/usecase"
	"shop-service/package/response"
	"shop-service/package/structs"

	"github.com/labstack/echo/v4"
)

type ShopHTTP struct{
	uc usecase.IShop
}
func (h ShopHTTP) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	resp, err := h.uc.GetAll(ctx); 
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success get all shop")
}


// Create 	godoc
// @Tags  	auth
// @Status  200  {object}  structs.Response
// @Router  /shops [post]
func (h ShopHTTP) Create(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestCreateShop{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.Create(ctx, req); err != nil {
		return err
	}
	return response.JSONSuccess(c, nil, "success create shop")
}

func NewShopHTTP(r *echo.Group, uc usecase.IShop)  {
	u := ShopHTTP{uc: uc}
	routeGroup := r.Group("/shops")
	routeGroup.GET("", u.GetAll).Name = "shop.get-all"
	routeGroup.POST("", u.Create).Name = "shop.create"
}