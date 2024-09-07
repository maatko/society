package view

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template/auth"
)

func LogIntoAccount(writer http.ResponseWriter, request *http.Request) {
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

	session, err := model.NewSession(user, 12*time.Hour)
	if err != nil {
		http.Redirect(writer, request, "/register", http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "session",
		Value:    session.UUID.String(),
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func ShowLogin(writer http.ResponseWriter, request *http.Request) {
	auth.Login().Render(request.Context(), writer)
}
