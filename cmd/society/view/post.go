package view

import (
	"log"
	"net/http"
	"strconv"

	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template"
	"github.com/maatko/society/web/template/component"
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

func POST_LikePost(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		return
	}

	err = request.ParseForm()
	if err != nil {
		return
	}

	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		return
	}

	post, err := model.GetPostByID(id)
	if err != nil {
		return
	}

	err = post.Like(user)
	if err != nil {
		return
	}

	component.Post(user, post).Render(request.Context(), writer)
}

func DELETE_Post(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	err = request.ParseForm()
	if err != nil {
		return
	}

	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		return
	}

	post, err := model.GetPostByID(id)
	if err != nil {
		return
	}

	if user.ID != post.User.ID {
		return
	}

	post.Delete()
}
