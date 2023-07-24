package server

import (
	"TestTaskCodefinity/internal/handler"
	"TestTaskCodefinity/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type App struct {
	app *fiber.App
}

func NewApp() *App {
	return &App{
		app: fiber.New(),
	}
}
func (a *App) Run() {
	api := a.app.Group("/api")

	fizzBuzzHandler := handler.NewFizzBuzzHandler(service.NewFizzBuzzService())
	api.Post("/fizzbuzz", fizzBuzzHandler.GetStringArr)

	if err := a.app.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}

func (a *App) Shutdown() {
	if err := a.app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
