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

func GetPostByID(id int) (*Post, error) {
	row := server.DataBase().QueryRow("SELECT uuid, user, cover, about, created_at, updated_at FROM post WHERE id=?", id)

	var uuid uuid.UUID
	var userID int
	var cover string
	var about string
	var created_at time.Time
	var updated_at time.Time

	if err := row.Scan(&uuid, &userID, &cover, &about, &created_at, &updated_at); err != nil {
		return nil, err
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return &Post{
		ID:        id,
		UUID:      uuid,
		User:      user,
		Cover:     cover,
		About:     about,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
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

	rows, err := server.DataBase().Query("SELECT * FROM post ORDER BY created_at DESC")
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

func SearchPosts(query string) []*Post {
	posts := []*Post{}

	rows, err := server.DataBase().Query("SELECT * FROM post WHERE about LIKE '%" + query + "%'")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		post := &Post{}

		var userID int
		if err := rows.Scan(&post.ID, &userID, &post.UUID, &post.Cover, &post.About, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil
		}

		user, err := GetUserByID(userID)
		if err != nil {
			return nil
		}

		post.User = user
		posts = append(posts, post)
	}

	return posts
}

func (post *Post) Like(user *User) error {
	var err error
	if post.IsLikedBy(user) {
		_, err = server.DataBase().Exec("DELETE FROM like WHERE user=? AND post=?", user.ID, post.ID)
	} else {
		_, err = server.DataBase().Exec("INSERT INTO like (user, post) VALUES (?, ?)", user.ID, post.ID)
	}
	return err
}

func (post *Post) IsLikedBy(user *User) bool {
	row := server.DataBase().QueryRow("SELECT COUNT(*) FROM like WHERE user=? AND post=?", user.ID, post.ID)

	var likeCount int
	if err := row.Scan(&likeCount); err != nil {
		return false
	}

	return likeCount > 0
}

func (post *Post) GetLikes() []*Like {
	posts := []*Like{}

	rows, err := server.DataBase().Query("SELECT id, user FROM like WHERE post=?", post.ID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		like := &Like{
			Post: post,
		}

		var userID int
		if err := rows.Scan(&like.ID, &userID); err != nil {
			return nil
		}

		user, err := GetUserByID(userID)
		if err != nil {
			return nil
		}

		like.User = user
		posts = append(posts, like)
	}

	return posts
}
