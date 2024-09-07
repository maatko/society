package view

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template/auth"
)

func RegisterAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "/register", http.StatusBadRequest)
		return
	}

	name := request.FormValue("name")
	password := request.FormValue("password")

	_, err = model.GetUserByName(name)
	if err == nil {
		http.Redirect(writer, request, "/register", http.StatusTemporaryRedirect)
		return
	}

	user, err := model.NewUser(name, password)
	if err != nil {
		http.Redirect(writer, request, "/register", http.StatusInternalServerError)
		return
	}

	err = Login(user, 12*time.Hour, writer)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func ShowRegister(writer http.ResponseWriter, request *http.Request) {
	auth.Register().Render(request.Context(), writer)
}
