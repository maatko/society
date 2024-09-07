package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
)

func LogoutOfAccount(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("session")
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	session, err := model.GetSessionByCookie(cookie)
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusInternalServerError)
		return
	}

	err = session.Delete()
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusInternalServerError)
		return
	}

	server.DeleteCookie(writer, "session")
	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
