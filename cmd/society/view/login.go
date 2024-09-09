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
		auth.Login("internal server error").Render(request.Context(), writer)
		return
	}

	name := request.FormValue("name")
	password := request.FormValue("password")

	user, err := model.GetUser(name, password)
	if err != nil {
		auth.Login("invalid credentials").Render(request.Context(), writer)
		return
	}

	err = authentication.Login(writer, user)
	if err != nil {
		auth.Login("failed to authenticate").Render(request.Context(), writer)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func GET_Login(writer http.ResponseWriter, request *http.Request) {
	auth.Login("").Render(request.Context(), writer)
}
