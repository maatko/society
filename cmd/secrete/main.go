package main

import (
	"log"
	"net/http"

	"github.com/maatko/secrete/api"
	"github.com/maatko/secrete/web/template"
	"github.com/maatko/secrete/web/template/auth"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	template.Index().Render(request.Context(), writer)
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	auth.Login().Render(request.Context(), writer)
}

func main() {
	db, err := api.NewDataBase("./api/model/", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)

	// static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe(":8080", nil)

	db.Close()
}
