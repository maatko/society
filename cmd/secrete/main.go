package main

import (
	"net/http"

	"github.com/maatko/secrete/web/template"
)

func BasicHandler(writer http.ResponseWriter, request *http.Request) {
	template.Index("secrete").Render(request.Context(), writer)
}

func main() {
	http.HandleFunc("/", BasicHandler)
	http.ListenAndServe(":8080", nil)
}
