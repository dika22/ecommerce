package delivery

import (
	"auth-service/internal/domain/user/usecase"
	"auth-service/package/response"
	"auth-service/package/structs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

type UserHTTP struct {
	uc usecase.IUser
}

func (h UserHTTP) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestSignUp{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	if err := h.uc.SignUp(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success create user")
}

func (h UserHTTP) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestLogin{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	resp, err := h.uc.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Str("endpoint", "/login").Msg("not authorized")
		return response.JSONResponse(c, http.StatusUnauthorized, false, err.Error(), nil)
	}
	return response.JSONResponse(c, http.StatusOK, true, "success login", resp)
}

func (h UserHTTP) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	req := structs.RequestUpdateUser{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	req.ID = cast.ToInt64(id)
	if err := h.uc.UpdateUser(ctx, req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, false, err.Error(), nil)
	}
	return response.JSONSuccess(c, nil, "success update user")
}

func NewUserHTTP(r *echo.Group, uc usecase.IUser) {
	u := UserHTTP{uc: uc}
	// r.Group("/users")
	r.POST("/signup", u.SignUp).Name = "users.signup"
	r.POST("/login", u.Login).Name = "users.login"
	r.PUT("/update/:id", u.UpdateUser).Name = "update.user"
}
