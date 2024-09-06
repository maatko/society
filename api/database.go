package api

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/maatko/secrete/api/model"
	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	Instance *sql.DB
}

func NewDataBase(modelDirectory string, connection string) (*DataBase, error) {
	instance, err := sql.Open("sqlite3", connection)
	if err != nil {
		return nil, err
	}

	filepath.Walk(modelDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".sql") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = instance.Exec(string(data))
			if err != nil {
				log.Fatal(err)
			}

			log.Print("Model Registered: ", info.Name())
		}
		return nil
	})

	return &DataBase{
		Instance: instance,
	}, nil
}

func (db *DataBase) CreateUser(name string, password string) (*model.User, error) {
	_, err := db.Instance.Exec("INSERT INTO user (name, password) VALUES (?, ?)", name, password)
	if err != nil {
		return nil, err
	}

	return db.GetUser(name, password)
}

func (db *DataBase) DeleteUser(id int) error {
	_, err := db.Instance.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) DeleteUserByCredentials(name string, password string) error {
	user, err := db.GetUser(name, password)
	if err != nil {
		return err
	}
	return db.DeleteUser(user.ID)
}

func (db *DataBase) GetUser(name string, password string) (*model.User, error) {
	row := db.Instance.QueryRow("SELECT id FROM user WHERE name=? AND password=?", name, password)

	var id int
	if err := row.Scan(&id); err != nil {
		return nil, err
	}

	return model.NewUser(id, name, password), nil
}

func (db *DataBase) GetUserByName(name string) (*model.User, error) {
	row := db.Instance.QueryRow("SELECT id, password FROM user WHERE name=?", name)

	var id int
	var password string
	if err := row.Scan(&id, &password); err != nil {
		return nil, err
	}

	return model.NewUser(id, name, password), nil
}

func (db *DataBase) CreateSession(user *model.User, duration time.Duration) (*model.Session, error) {
	created_at := time.Now()
	expires_at := time.Now().Add(duration)

	_, err := db.Instance.Exec("INSERT INTO session (user, created_at, expires_at) VALUES (?, ?, ?)", user.ID, created_at, expires_at)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (db *DataBase) DeleteSession(user *model.User) error {
	_, err := db.Instance.Exec("DELETE FROM session WHERE user=?", user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) DeleteExpiredSessions(time time.Time) error {
	_, err := db.Instance.Exec("DELETE FROM session WHERE expires_at<=?", time)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) GetSession(user *model.User) (*model.Session, error) {
	row := db.Instance.QueryRow("SELECT id, created_at, expires_at FROM session WHERE user=?", user.ID)

	var id int
	var created_at time.Time
	var expires_at time.Time

	if err := row.Scan(&id, &created_at, &expires_at); err != nil {
		return nil, err
	}

	return model.NewSession(id, user, created_at, expires_at), nil
}

func (db *DataBase) Close() {
	db.Instance.Close()
	db.Instance = nil
}
