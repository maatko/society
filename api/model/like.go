package model

import (
	"github.com/maatko/society/internal/server"
)

type Like struct {
	ID   int
	User *User
	Post *Post
}

func NewLike(user *User, post *Post) (*Like, error) {
	result, err := server.DataBase().Exec("INSERT INTO like (user, post) VALUES (?, ?)", user.ID, post.ID)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetLikeByID(int(id))
}

func GetLikeByID(id int) (*Like, error) {
	row := server.DataBase().QueryRow("SELECT user, post FROM like WHERE id=?", id)

	var userID int
	var postID int

	if err := row.Scan(&userID, &postID); err != nil {
		return nil, err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	post, err := GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	return &Like{
		ID:   id,
		User: user,
		Post: post,
	}, nil
}
