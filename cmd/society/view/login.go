package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	authentication "github.com/maatko/society/internal/auth"
	"github.com/maatko/society/web/template/auth"
)

func POST_Login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "/register", http.StatusBadRequest)
		return
	}

	name := request.FormValue("name")
	password := request.FormValue("password")

	user, err := model.GetUser(name, password)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	err = authentication.Login(writer, user, 43200)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func GET_Login(writer http.ResponseWriter, request *http.Request) {
	auth.Login().Render(request.Context(), writer)
}
