package model

type User struct {
	ID       int
	Name     string
	Password string
}

func NewUser(id int, name string, password string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Password: password,
	}
}
