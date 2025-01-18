package http

import (
	"net/http"

	"github.com/Satria2133/echo-CRUD/internal/auth/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(e *echo.Echo, uc usecase.AuthUsecase) {
	handler := &AuthHandler{usecase: uc}
	e.POST("/login", handler.Login)
}

func (h *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := h.usecase.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
