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

	server.AddRoute("/", view.GET_Home)

	// authentication routes
	server.AddRoute("GET  /login", view.GET_Login)
	server.AddRoute("POST /login", view.POST_Login)

	server.AddRoute("GET  /register", view.GET_Register)
	server.AddRoute("POST /register", view.POST_Register)

	server.AddRoute("GET  /logout", view.GET_Logout)

	server.Start(":8080", middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
