package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func GET_Home(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusInternalServerError)
		return
	}

	template.Home(user).Render(request.Context(), writer)
}
