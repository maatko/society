package main

import (
	"log"

	"github.com/maatko/society/cmd/society/view"
	"github.com/maatko/society/internal/middleware"
	"github.com/maatko/society/internal/server"
)

func main() {
	err := server.Setup("./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	server.AddRoute("/", view.ShowHome)
	server.AddRoute("/login", view.ShowLogin)
	server.AddRoute("/register", view.ShowRegister)

	server.Start(":8080", middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
