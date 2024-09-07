package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	DataBase *sql.DB
	Router   *http.ServeMux
}

type MiddlewareCallback func(http.Handler) http.Handler

var (
	Instance *Server
)

func Setup(connection string) error {
	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		return err
	}

	filepath.Walk("./api/model", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if strings.HasSuffix(path, ".sql") {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			_, err = db.Exec(string(data))
			if err != nil {
				log.Fatal(err)
			}

			log.Print("Model Registered: ", info.Name())
		}
		return nil
	})

	Instance = &Server{
		DataBase: db,
		Router:   http.NewServeMux(),
	}

	return nil
}

func Start(address string, middlewares ...MiddlewareCallback) {
	// static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	Instance.Router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// this function applies all the middlewares
	// that were provided
	apply := func(handler http.Handler, middlewares ...MiddlewareCallback) http.Handler {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		return handler
	}

	http.ListenAndServe(address, apply(Instance.Router, middlewares...))

	// make sure to close the database
	// once the server has stopped
	Instance.DataBase.Close()
}

func AddRoute(route string, handler func(http.ResponseWriter, *http.Request)) {
	Instance.Router.HandleFunc(route, handler)
}

func SetCookie(writer http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(writer, cookie)
}

func DeleteCookie(writer http.ResponseWriter, name string) {
	SetCookie(writer, &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	})
}

func DataBase() *sql.DB {
	return Instance.DataBase
}
