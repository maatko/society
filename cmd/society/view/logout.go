package view

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
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

	// remove the cookie
	http.SetCookie(writer, &http.Cookie{
		Name:    "session",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
