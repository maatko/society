package middleware

import (
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		path := request.RequestURI

		cookie, err := request.Cookie("session")
		if err != nil {
			if strings.HasPrefix(path, "/login") || strings.HasPrefix(path, "/register") || strings.HasPrefix(path, "/static") {
				next.ServeHTTP(writer, request)
				return
			}

			http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
			next.ServeHTTP(writer, request)
			return
		}

		log.Println(cookie.Value)

		if strings.HasPrefix(path, "/login") || strings.HasPrefix(path, "/register") || strings.HasPrefix(path, "/static") {
			http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
			next.ServeHTTP(writer, request)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
