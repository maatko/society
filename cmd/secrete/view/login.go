package view

import (
	"net/http"

	"github.com/maatko/secrete/internal/server"
	"github.com/maatko/secrete/web/template/auth"
)

func ShowLogin(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Login())
}
