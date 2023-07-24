package main

import "TestTaskCodefinity/pkg/server"

func main() {
	app := server.NewApp()
	app.Run()
	defer app.Shutdown()
}
