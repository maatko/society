package main

import (
	"log"
	"net/http"

	"github.com/maatko/secrete/api"
	"github.com/maatko/secrete/web/template"
)

func BasicHandler(writer http.ResponseWriter, request *http.Request) {
	template.Index().Render(request.Context(), writer)
}

func main() {
	db, err := api.NewDataBase("./api/model/", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", BasicHandler)
	http.ListenAndServe(":8080", nil)

	db.Close()
}
