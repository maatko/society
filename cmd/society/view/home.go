package view

import (
	"log"
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"
)

func GET_Home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		template.NotFound().Render(request.Context(), writer)
		return
	}

	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	posts, err := model.GetAllPosts()
	if err != nil {
		log.Println(err)
		posts = []*model.Post{}
	}

	template.Home(user, posts).Render(request.Context(), writer)
}
