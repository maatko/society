package view

import (
	"log"
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
		log.Println(err)
		http.Redirect(writer, request, "/register", http.StatusTemporaryRedirect)
		return
	}

	user, err := model.NewUser(name, password)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/register", http.StatusInternalServerError)
		return
	}

	err = authentication.Login(writer, user, 43200)
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func GET_Register(writer http.ResponseWriter, request *http.Request) {
	auth.Register().Render(request.Context(), writer)
}
