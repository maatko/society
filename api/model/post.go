package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/maatko/society/internal/server"
)

type Post struct {
	ID        int
	User      *User
	UUID      uuid.UUID
	Cover     string
	About     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(user *User, uuid uuid.UUID, cover string, about string) (*Post, error) {
	_, err := server.DataBase().Exec("INSERT INTO post (user, uuid, cover, about) VALUES (?, ?, ?, ?)", user.ID, uuid.String(), cover, about)
	if err != nil {
		return nil, err
	}

	return GetPostByUUID(uuid)
}

func GetPostByUUID(uuid uuid.UUID) (*Post, error) {
	row := server.DataBase().QueryRow("SELECT id, user, cover, about, created_at, updated_at FROM post WHERE uuid=?", uuid.String())

	var id int
	var userID int
	var cover string
	var about string
	var created_at time.Time
	var updated_at time.Time

	if err := row.Scan(&id, &userID, &cover, &about, &created_at, &updated_at); err != nil {
		return nil, err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:        id,
		User:      user,
		Cover:     cover,
		About:     about,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func GetAllPosts() ([]*Post, error) {
	posts := []*Post{}

	rows, err := server.DataBase().Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &Post{}

		var userID int
		if err := rows.Scan(&post.ID, &userID, &post.UUID, &post.Cover, &post.About, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}

		user, err := GetUserByID(userID)
		if err != nil {
			return nil, err
		}

		post.User = user
		posts = append(posts, post)
	}

	return posts, nil
}
