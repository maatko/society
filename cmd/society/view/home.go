package view

import (
	"net/http"

	"github.com/maatko/secrete/internal/server"
	"github.com/maatko/secrete/web/template"
)

func ShowHome(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, template.Index())
}
