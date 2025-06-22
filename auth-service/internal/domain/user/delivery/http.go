package delivery

import (
	"auth-service/internal/domain/user/usecase"
	"auth-service/package/response"
	"auth-service/package/structs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UserHTTP struct {
	uc usecase.IUser
}

func (h UserHTTP) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestSignUp{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	if err := h.uc.SignUp(ctx, req); err != nil {
		return err
	}
	return response.JSONSuccess(c, nil, "success create user")
}

func (h UserHTTP) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := structs.RequestLogin{}
	if err := c.Bind(&req); err != nil {
		return response.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
	}
	resp, err := h.uc.Login(ctx, req)
	if err != nil {
		log.Error().Err(err).Str("endpoint", "/login").Msg("not authorized")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return response.JSONResponse(c, http.StatusOK, "success", "success login", resp)
}

func NewUserHTTP(r *echo.Group, uc usecase.IUser) {
	u := UserHTTP{uc: uc}
	r.POST("/signup", u.SignUp).Name = "users.signup"
	r.POST("/login", u.Login).Name = "users.login"
}
