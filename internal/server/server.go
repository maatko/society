package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/maatko/secrete/api"
)

type Server struct {
	DataBase *api.DataBase
	Router   *http.ServeMux
}

type MiddlewareCallback func(http.Handler) http.Handler

var (
	Instance *Server
)

func Setup(connection string) error {
	db, err := api.NewDataBase("./api/model/", connection)
	if err != nil {
		return err
	}

	Instance = &Server{
		DataBase: db,
		Router:   http.NewServeMux(),
	}

	return nil
}

func Start(address string, middlewares ...MiddlewareCallback) {
	// static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

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

func Render(writer http.ResponseWriter, request *http.Request, component templ.Component) error {
	return component.Render(request.Context(), writer)
}

func DataBase() *api.DataBase {
	return Instance.DataBase
}
