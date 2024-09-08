package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	authentication "github.com/maatko/society/internal/auth"
	"github.com/maatko/society/web/template/auth"
)

func POST_Register(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "/register", http.StatusBadRequest)
		return
	}

	name := request.FormValue("name")
	password := request.FormValue("password")

	_, err = model.GetUserByName(name)
	if err == nil {
		auth.Register("user already exists!").Render(request.Context(), writer)
		return
	}

	user, err := model.NewUser(name, password)
	if err != nil {
		auth.Register("internal server error!").Render(request.Context(), writer)
		return
	}

	err = authentication.Login(writer, user, 43200)
	if err != nil {
		auth.Register("failed to authenticate").Render(request.Context(), writer)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func GET_Register(writer http.ResponseWriter, request *http.Request) {
	auth.Register("").Render(request.Context(), writer)
}
