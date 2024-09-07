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

	// authentication routes
	server.AddRoute("GET  /login", view.ShowLogin)
	server.AddRoute("POST /login", view.LogIntoAccount)

	server.AddRoute("GET  /register", view.ShowRegister)
	server.AddRoute("POST /register", view.RegisterAccount)

	server.AddRoute("GET  /logout", view.LogoutOfAccount)

	server.Start(":8080", middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
