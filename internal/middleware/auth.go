package middleware

import (
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		path := request.RequestURI

		if strings.HasPrefix(path, "/login") || strings.HasPrefix(path, "/register") || strings.HasPrefix(path, "/static") {
			next.ServeHTTP(writer, request)
			return
		}

		_, err := request.Cookie("session")
		if err != nil {
			http.Redirect(writer, request, "/login", http.StatusPermanentRedirect)
			next.ServeHTTP(writer, request)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
