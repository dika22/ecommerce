package delivery

import (
	"auth-service/internal/domain/seller/usecase"
	"auth-service/package/response"
	"auth-service/package/structs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)


type SellerHTTP struct {
	uc usecase.ISeller
}

func (s SellerHTTP) CreateSeller(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestCreateSeller{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	if err := s.uc.CreateSeller(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success create user")
}

func (s SellerHTTP) UpdateSeller(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := structs.RequestUpdateSeller{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}

	req.Id = cast.ToInt64(id)
	if err := s.uc.UpdateSeller(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success update user")
}


func NewSellerHTTP(r *echo.Group, uc usecase.ISeller) {
	u := SellerHTTP{uc: uc}
	r.POST("/sellers", u.CreateSeller).Name = "create.seller"
	r.PUT("/sellers/:id", u.UpdateSeller).Name = "update.seller"
}
