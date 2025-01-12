package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/networking-tic-tac-toe/handler"
	"github.com/toastsandwich/networking-tic-tac-toe/repository"
	"github.com/toastsandwich/networking-tic-tac-toe/service"
	"github.com/toastsandwich/networking-tic-tac-toe/validator"
	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("./db/database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	validator := validator.NewCustomValidator()

	e := echo.New() // Echo
	e.Use(
		middleware.Logger(),
		middleware.CORS(),
		// middleware.CSRF(),
	)
	e.Validator = validator // Validator

	api := e.Group("/api")

	api.POST("/login", handler.LoginHandler)
	api.POST("/logout", handler.LogoutHandler)
	user := api.Group("/user")
	user.GET("/get", handler.GetUserHandler)
	user.POST("/create", handler.CreateUserHandler)
	user.DELETE("/delete", handler.DeleteUserHandler)

	if err := e.Start(":3001"); err != nil {
		e.Logger.Fatal(err)
	}
}
