package model

import (
	"time"

	"github.com/maatko/secrete/api/model"
	"github.com/maatko/secrete/internal/server"
)

type Session struct {
	ID        int
	User      *User
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewSession(user *User, duration time.Duration) (*Session, error) {
	created_at := time.Now()
	expires_at := time.Now().Add(duration)

	_, err := server.DataBase().Exec("INSERT INTO session (user, created_at, expires_at) VALUES (?, ?, ?)", user.ID, created_at, expires_at)
	if err != nil {
		return nil, err
	}

	return GetSession(user)
}

func GetSession(user *model.User) (*Session, error) {
	row := server.DataBase().QueryRow("SELECT id, created_at, expires_at FROM session WHERE user=?", user.ID)

	var id int
	var created_at time.Time
	var expires_at time.Time

	if err := row.Scan(&id, &created_at, &expires_at); err != nil {
		return nil, err
	}

	return &Session{
		ID:        id,
		User:      user,
		CreatedAt: created_at,
		ExpiresAt: expires_at,
	}, nil
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
