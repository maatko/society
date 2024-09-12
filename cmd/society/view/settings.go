package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func GET_Settings(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	template.Settings(user).Render(request.Context(), writer)
}
