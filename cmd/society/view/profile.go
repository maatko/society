package view

import (
	"net/http"
	"strings"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func ShowProfile(writer http.ResponseWriter, request *http.Request) bool {
	name := strings.TrimSuffix(request.URL.Path[1:], "/")

	profileUser, err := model.GetUserByName(name)
	if err != nil {
		return false
	}

	loggedInUser, err := model.GetUserByRequest(request)
	if err != nil {
		return false
	}

	posts, err := profileUser.GetPosts()
	if err != nil {
		return false
	}

	template.Profile(loggedInUser, profileUser, posts).Render(request.Context(), writer)
	return true
}
