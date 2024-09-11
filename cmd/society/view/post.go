package view

import (
	"log"
	"net/http"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template"
)

func GET_CreatePost(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	template.Post(user, "").Render(request.Context(), writer)
}

func POST_CreatePost(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	err = request.ParseMultipartForm(server.UPLOAD_LIMIT)
	if err != nil {
		template.Post(user, "invalid form data").Render(request.Context(), writer)
		return
	}

	imageID, imagePath, err := server.UploadImage(request, "image")
	if err != nil {
		template.Post(user, "failed to upload image").Render(request.Context(), writer)
		return
	}

	_, err = model.NewPost(user, imageID, imagePath, request.FormValue("about"))
	if err != nil {
		log.Println(err)
		template.Post(user, "failed to create post in database").Render(request.Context(), writer)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
