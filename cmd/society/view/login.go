package view

import (
	"net/http"

	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template/auth"
)

func ShowLogin(writer http.ResponseWriter, request *http.Request) {
	server.Render(writer, request, auth.Login())
}
