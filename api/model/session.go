package model

import "time"

type Session struct {
	ID        int
	User      *User
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewSession(id int, user *User, createdAt time.Time, expiresAt time.Time) *Session {
	return &Session{
		ID:        id,
		User:      user,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}
