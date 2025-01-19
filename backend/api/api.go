package main

import (
	"log"
	"runtime"
	"runtime/debug"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/networking-tic-tac-toe/handler"
	matchmaker "github.com/toastsandwich/networking-tic-tac-toe/match_maker"
	"github.com/toastsandwich/networking-tic-tac-toe/repository"
	"github.com/toastsandwich/networking-tic-tac-toe/service"
	"github.com/toastsandwich/networking-tic-tac-toe/validator"
	"go.etcd.io/bbolt"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // this will set maximum number of go routines
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	db, err := bbolt.Open("./db/database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	repository, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal(err)
	}
	// all necessary stuff
	matchmaker := matchmaker.NewMatchMaker()
	service := service.NewService(repository)
	handler := handler.NewHandler(service, matchmaker)
	validator := validator.NewCustomValidator()
	gameServer := NewGameServer("localhost", "3002", service)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := gameServer.Start(); err != nil {
			log.Fatal(err)
		}
	}(wg)
	// echo
	e := echo.New() // Echo
	loggerConfig := middleware.LoggerConfig{
		Format: "protocol=${protocol}, method=${method}, uri=${uri}, status=${status}, error=${error}, latency=${latency_human}\n",
	}
	e.Use(
		middleware.LoggerWithConfig(loggerConfig),
		middleware.CORS(),
	)
	e.Validator = validator // Validator

	api := e.Group("/api")

	api.POST("/login", handler.LoginHandler)
	api.POST("/logout", handler.LogoutHandler, handler.AuthMiddleware)
	user := api.Group("/user")
	user.GET("/get", handler.GetUserHandler)
	user.POST("/create", handler.CreateUserHandler)
	user.DELETE("/delete", handler.DeleteUserHandler)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if err := e.Start(":3001"); err != nil {
			e.Logger.Fatal(err)
			wg.Done()
		}
	}(wg)
	wg.Wait()
}
