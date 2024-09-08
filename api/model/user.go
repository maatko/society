package model

import (
	"net/http"

	"github.com/maatko/society/internal/server"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Name     string
	Password string
}

func NewUser(name string, password string) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = server.DataBase().Exec("INSERT INTO user (name, password) VALUES (?, ?)", name, passwordHash)
	if err != nil {
		return nil, err
	}

	return GetUser(name, password)
}

func GetUser(name string, password string) (*User, error) {
	user, err := GetUserByName(name)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
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
