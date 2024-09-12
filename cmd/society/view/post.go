package view

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/maatko/society/api/model"
	"github.com/maatko/society/internal/server"
	"github.com/maatko/society/web/template/component"
	"github.com/maatko/society/web/template/post"
)

func GET_CreatePost(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	id := strings.TrimSuffix(request.URL.Path[6:], "/")
	if len(id) > 0 {
		uuid, err := uuid.Parse(id)
		if err != nil {
			http.Redirect(writer, request, "/post", http.StatusTemporaryRedirect)
			return
		}

		userPost, err := model.GetPostByUUID(uuid)
		if err != nil {
			http.Redirect(writer, request, "/post", http.StatusTemporaryRedirect)
			return
		}

		post.ViewPost(user, userPost).Render(request.Context(), writer)
		return
	}

	post.CreatePost(user, "").Render(request.Context(), writer)
}

func POST_CreatePost(writer http.ResponseWriter, request *http.Request) {
	user, err := model.GetUserByRequest(request)
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusTemporaryRedirect)
		return
	}

	err = request.ParseMultipartForm(server.UPLOAD_LIMIT)
	if err != nil {
		post.CreatePost(user, "invalid form data").Render(request.Context(), writer)
		return
	}

	imageID, imagePath, err := server.UploadImage(request, "image")
	if err != nil {
		post.CreatePost(user, "failed to upload image").Render(request.Context(), writer)
		return
	}

	_, err = model.NewPost(user, imageID, imagePath, request.FormValue("about"))
	if err != nil {
		log.Println(err)
		post.CreatePost(user, "failed to create post in database").Render(request.Context(), writer)
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
