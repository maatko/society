package view

import (
	"net/http"

	"github.com/maatko/society/web/template"
)

func ShowHome(writer http.ResponseWriter, request *http.Request) {
	template.Home().Render(request.Context(), writer)
}
