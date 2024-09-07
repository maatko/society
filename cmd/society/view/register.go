package view

import (
	"net/http"

	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template/auth"
)

func ShowRegister(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Register())
}
