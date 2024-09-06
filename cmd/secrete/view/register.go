package view

import (
	"net/http"

	"github.com/maatko/secrete/internal/server"
	"github.com/maatko/secrete/web/template/auth"
)

func ShowRegister(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Register())
}
