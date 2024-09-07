package model

import (
	"net/http"

	"github.com/maatko/society/internal/server"
)

type User struct {
	ID       int
	Name     string
	Password string
}

func NewUser(name string, password string) (*User, error) {
	_, err := server.DataBase().Exec("INSERT INTO user (name, password) VALUES (?, ?)", name, password)
	if err != nil {
		return nil, err
	}

	return GetUser(name, password)
}

func GetUser(name string, password string) (*User, error) {
	row := server.DataBase().QueryRow("SELECT id FROM user WHERE name=? AND password=?", name, password)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}, nil
}

func GetUserByID(id int) (*User, error) {
	row := server.DataBase().QueryRow("SELECT name, password FROM user WHERE id=?", id)

	var name string
	var password string

	if err := row.Scan(&name, &password); err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}, nil
}

func GetUserByName(name string) (*User, error) {
	row := server.DataBase().QueryRow("SELECT id, password FROM user WHERE name=?", name)

	var id int
	var password string

	if err := row.Scan(&id, &password); err != nil {
		return nil, err
	}

	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}, nil
}

func GetUserByRequest(request *http.Request) (*User, error) {
	cookie, err := request.Cookie("session")
	if err != nil {
		return nil, err
	}

	session, err := GetSessionByCookie(cookie)
	if err != nil {
		return nil, err
	}

	return session.User, nil
}

func (user *User) Delete() error {
	_, err := server.DataBase().Exec("DELETE FROM user WHERE id=?", user.ID)
	if err != nil {
		return err
	}
	return nil
}
