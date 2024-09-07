package view

import (
	"net/http"

	"github.com/maatko/society/internal/auth"
)

func GET_Logout(writer http.ResponseWriter, request *http.Request) {
	err := auth.Logout(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/", http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
