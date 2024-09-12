package view

import (
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func GET_Search(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	query := request.URL.Query().Get("query")

	template.Search(user, model.SearchUsers(query), model.SearchPosts(query)).Render(request.Context(), writer)
}
