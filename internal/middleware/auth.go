package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		path := request.RequestURI

		_, err := request.Cookie("session")
		if err != nil {
			if strings.HasPrefix(path, "/login") || strings.HasPrefix(path, "/register") || strings.HasPrefix(path, "/static") {
				next.ServeHTTP(writer, request)
				return
			}

			http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
			next.ServeHTTP(writer, request)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
