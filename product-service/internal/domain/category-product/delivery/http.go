package delivery

import (
	"context"
	"net/http"
	"product-service/internal/domain/category-product/usecase"
	"product-service/package/response"
	"product-service/package/structs"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type CategoryProductHTTP struct {
	uc usecase.ICategoryProduct
}

func (u CategoryProductHTTP) GetAll(c echo.Context) error {
	ctx:= context.Background()
	resp, err := u.uc.GetAll(ctx)
	if err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, resp, "success")
}

func (u CategoryProductHTTP) CreateCategoryProduct(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestCreateCategoryProduct{}
	if err := c.Bind(req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	if err := u.uc.Create(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONResponse(c, http.StatusOK, true, "success", nil)
	
}


func (u CategoryProductHTTP) UpdateCategoryProduct(c echo.Context) error {
	ctx := c.Request().Context()
	req := &structs.RequestUpdateCategoryProduct{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := u.uc.Update(ctx, req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

func (u CategoryProductHTTP) DeleteCategoryProduct(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	if err := u.uc.Delete(ctx, cast.ToInt64(id)); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")
}

func NewCategoryProductHTTP(r *echo.Group, uc usecase.ICategoryProduct)  {
	u := CategoryProductHTTP{uc: uc}
	r.Group("/category-product")
	r.GET("", u.GetAll).Name = "category-product.get-all"
	r.POST("", u.CreateCategoryProduct).Name = "category-product.create"
	r.PUT("/:id", u.UpdateCategoryProduct).Name = "category-product.update"
	r.DELETE("/:id", u.DeleteCategoryProduct).Name = "category-product.delete"
}