package api

import (
	"database/sql"
	"os"
	"path/filepath"
	"strings"

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
				return err
			}
		}
		return nil
	})

	return &DataBase{
		Instance: instance,
	}, nil
}

func (db *DataBase) Close() {
	db.Instance.Close()
	db.Instance = nil
}
