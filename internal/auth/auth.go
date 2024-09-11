package auth

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
)

// 43200

const (
	SESSION_COOKIE      = "session"
	SESSION_COOKIE_PATH = "/"
	SESSION_DURATION    = 43200
)

func Login(writer http.ResponseWriter, user *model.User) error {
	session, err := model.NewSession(user, SESSION_DURATION*time.Second)
	if err != nil {
		return err
	}

	server.SetCookie(writer, &http.Cookie{
		Name:     SESSION_COOKIE,
		Value:    session.UUID.String(),
		Path:     SESSION_COOKIE_PATH,
		MaxAge:   int(SESSION_DURATION),
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func Logout(writer http.ResponseWriter, request *http.Request) error {
	cookie, err := request.Cookie("session")
	if err != nil {
		return err
	}

	session, err := model.GetSessionByCookie(cookie)
	if err != nil {
		return err
	}

	err = session.Delete()
	if err != nil {
		return err
	}

	server.DeleteCookie(writer, SESSION_COOKIE)
	return nil
}
