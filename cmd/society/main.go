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

	// authentication routes
	server.AddRoute("GET  /login", view.GET_Login)
	server.AddRoute("POST /login", view.POST_Login)
	server.AddRoute("GET  /register", view.GET_Register)
	server.AddRoute("POST /register", view.POST_Register)
	server.AddRoute("GET  /logout", view.GET_Logout)

	// post management routes
	server.AddRoute("GET /post/", view.GET_CreatePost)
	server.AddRoute("POST /post", view.POST_CreatePost)
	server.AddRoute("POST /comment", view.POST_CommentPost)
	server.AddRoute("POST /like", view.POST_LikePost)
	server.AddRoute("DELETE /post", view.DELETE_Post)

	// serach
	server.AddRoute("GET /search", view.GET_Search)

	// settings
	server.AddRoute("GET /settings", view.GET_Settings)

	// home page
	server.AddRoute("/", view.GET_Home)

	server.Start("192.168.1.9:8080", middleware.LoggingMiddleware, middleware.AuthMiddleware)
}
