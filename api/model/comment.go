package model

import (
	"fmt"
	"time"

	"github.com/maatko/society/internal/server"
)

type Comment struct {
	ID        int
	User      *User
	Post      *Post
	Text      string
	CreatedAt time.Time
}

func NewComment(user *User, post *Post, text string) (*Comment, error) {
	result, err := server.DataBase().Exec("INSERT INTO comment (user, post, text) VALUES (?, ?, ?)", user.ID, post.ID, text)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetCommentByID(int(id))
}

func GetCommentByID(id int) (*Comment, error) {
	row := server.DataBase().QueryRow("SELECT user, post, text, created_at FROM comment WHERE id=?", id)

	var userID int
	var postID int
	var text string
	var createdAt time.Time

	if err := row.Scan(&userID, &postID, &text, &createdAt); err != nil {
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

	return &Comment{
		ID:        id,
		User:      user,
		Post:      post,
		Text:      text,
		CreatedAt: createdAt,
	}, nil
}

func (comment *Comment) GetTimeSince(time time.Time) string {
	elapsed := time.Sub(comment.CreatedAt)

	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60

	format := ""
	if hours > 0 {
		format += fmt.Sprintf("%02dh ", hours)
	} else {
		if minutes > 0 {
			format += fmt.Sprintf("%02dm ", minutes)
		} else {
			format += fmt.Sprintf("%02ds ", seconds)
		}
	}

	return format
}
