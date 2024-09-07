package main

import (
	"log"

	"github.com/maatko/secrete/cmd/secrete/view"
	"github.com/maatko/secrete/internal/middleware"
	"github.com/maatko/secrete/internal/server"
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
