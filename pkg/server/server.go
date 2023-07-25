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

type Config struct {
	Address    string
	FizzNumber int
	BuzzNumber int
}

func NewApp() *App {
	return &App{
		app: fiber.New(),
	}
}

func (a *App) Run(conf Config) {
	fizzBuzzHandler := handler.NewFizzBuzzHandler(service.NewFizzBuzzSolver(conf.FizzNumber, conf.BuzzNumber))
	a.app.Post("/run", fizzBuzzHandler.Solve)

	if err := a.app.Listen(conf.Address); err != nil {
		log.Fatal(err)
	}
}

func (a *App) Shutdown() {
	if err := a.app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
