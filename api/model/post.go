package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        int
	User      User
	UUID      uuid.UUID
	Cover     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
