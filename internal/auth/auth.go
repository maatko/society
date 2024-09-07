package auth

import (
	"net/http"
	"time"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
)

const (
	SESSION_COOKIE      = "session"
	SESSION_COOKIE_PATH = "/"
)

func Login(writer http.ResponseWriter, user *model.User, duration time.Duration) error {
	session, err := model.NewSession(user, duration*time.Second)
	if err != nil {
		return err
	}

	server.SetCookie(writer, &http.Cookie{
		Name:     SESSION_COOKIE,
		Value:    session.UUID.String(),
		Path:     SESSION_COOKIE_PATH,
		MaxAge:   int(duration),
		Secure:   true,
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
