package main

import (
	"TestTaskCodefinity/pkg/server"
	"log"
	"os"
	"strconv"
)

func main() {
	app := server.NewApp()

	fizzNumber, err := strconv.Atoi(os.Getenv("FIZZ_NUMBER"))
	if err != nil {
		log.Fatalf("error occurs while convert string to int: %s", err)
	}

	buzzNumber, err := strconv.Atoi(os.Getenv("BUZZ_NUMBER"))
	if err != nil {
		log.Fatalf("error occurs while convert string to int: %s", err)
	}

	app.Run(server.Config{
		Address:    os.Getenv("ADDRESS"),
		FizzNumber: fizzNumber,
		BuzzNumber: buzzNumber,
	})

	defer app.Shutdown()
}
