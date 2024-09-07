package model

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/maatko/society/internal/server"
)

type Session struct {
	ID        int
	User      *User
	UUID      uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewSession(user *User, duration time.Duration) (*Session, error) {
	created_at := time.Now()
	expires_at := time.Now().Add(duration)

	_, err := server.DataBase().Exec("INSERT INTO session (user, uuid, created_at, expires_at) VALUES (?, ?, ?, ?)", user.ID, uuid.NewString(), created_at, expires_at)
	if err != nil {
		return nil, err
	}

	return GetSession(user)
}

func GetSession(user *User) (*Session, error) {
	row := server.DataBase().QueryRow("SELECT id, uuid, created_at, expires_at FROM session WHERE user=?", user.ID)

	var id int
	var uid uuid.UUID
	var created_at time.Time
	var expires_at time.Time

	if err := row.Scan(&id, &uid, &created_at, &expires_at); err != nil {
		return nil, err
	}

	return &Session{
		ID:        id,
		User:      user,
		UUID:      uid,
		CreatedAt: created_at,
		ExpiresAt: expires_at,
	}, nil
}

func GetSessionByUUID(uid string) (*Session, error) {
	row := server.DataBase().QueryRow("SELECT user FROM session WHERE uuid=?", uid)

	var userID int
	if err := row.Scan(&userID); err != nil {
		return nil, err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return GetSession(user)
}

func GetSessionByCookie(cookie *http.Cookie) (*Session, error) {
	return GetSessionByUUID(cookie.Value)
}

func DeleteExpiredSessions(time time.Time) error {
	_, err := server.DataBase().Exec("DELETE FROM session WHERE expires_at<=?", time)
	if err != nil {
		return err
	}
	return nil
}

func (session *Session) Delete() error {
	_, err := server.DataBase().Exec("DELETE FROM session WHERE user=?", session.User.ID)
	if err != nil {
		return err
	}
	return nil
}
