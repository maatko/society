package view

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template/auth"
)

func Login(user *model.User, duration time.Duration, writer http.ResponseWriter) error {
	session, err := model.NewSession(user, 43200*time.Second)
	if err != nil {
		return err
	}

	server.SetCookie(writer, &http.Cookie{
		Name:     "session",
		Value:    session.UUID.String(),
		Path:     "/",
		MaxAge:   43200,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

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

	err = Login(user, 12*time.Hour, writer)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func ShowLogin(writer http.ResponseWriter, request *http.Request) {
	auth.Login().Render(request.Context(), writer)
}
