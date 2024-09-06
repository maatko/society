package main

import (
	"log"
	"net/http"

	"github.com/maatko/secrete/internal/middleware"
	"github.com/maatko/secrete/internal/server"
	"github.com/maatko/secrete/web/template"
	"github.com/maatko/secrete/web/template/auth"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, template.Index())
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Login())
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Register())
}

func main() {
	err := server.Setup("./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	server.AddRoute("/", HomeHandler)
	server.AddRoute("/login", LoginHandler)
	server.AddRoute("/register", RegisterHandler)

	server.Start(":8080", middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
