package view

import (
	"log"
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func GET_Home(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello world!")

	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	template.Home(user).Render(request.Context(), writer)
}
