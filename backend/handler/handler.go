package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/networking-tic-tac-toe/model"
	"github.com/toastsandwich/networking-tic-tac-toe/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetUserHandler(c echo.Context) error {
	email := c.QueryParams().Get("email")
	if email != "" {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "email missing"})
	}
	user, err := h.Service.GetUserService(email)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUserHandler(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	user.Role = "user"
	return h.Service.CreateUserService(user)
}

func (h *Handler) DeleteUserHandler(c echo.Context) error {
	email := c.QueryParams().Get("email")
	return h.Service.DeleteUserService(email)
}

func (h *Handler) LoginHandler(c echo.Context) error {
	// payload for login handler
	payload := &struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}{}
	if err := c.Bind(payload); err != nil {
		return err
	}
	if err := c.Validate(payload); err != nil {
		return err
	}
	user, err := h.Service.GetUserService(payload.Email)
	if err != nil {
		return err
	}
	if payload.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Credentials"})
	}
	token, err := CreateJWTToken(payload.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"token":   token,
		"user":    user,
	})
}
