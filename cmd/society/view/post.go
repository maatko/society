package view

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/maatko/society/api/model"
	"github.com/maatko/society/web/template"

	"image/jpeg"
	"image/png"
)

const (
	// 10 MB limit
	FILE_SIZE_LIMIT  = 10 << 20
	STORAGE_DIR      = "./web/storage"
	FILE_PATH_FORMAT = STORAGE_DIR + "/%s.%s"
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

	err = request.ParseMultipartForm(FILE_SIZE_LIMIT)
	if err != nil {
		template.Post(user, "invalid form data").Render(request.Context(), writer)
		return
	}

	_, err = os.Stat(STORAGE_DIR)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(STORAGE_DIR, 0755)
			if err != nil {
				log.Println(err)
			}
		}
	}

	imageID, err := uuid.NewUUID()
	if err != nil {
		template.Post(user, "failed to generate new uuid").Render(request.Context(), writer)
		return
	}

	file, _, err := request.FormFile("image")
	if err != nil {
		template.Post(user, "failed to locate the image").Render(request.Context(), writer)
		return
	}
	defer file.Close()

	image, format, err := image.Decode(file)
	if err != nil {
		template.Post(user, "invalid image").Render(request.Context(), writer)
		return
	}

	if format != "png" && format != "jpg" && format != "jpeg" {
		template.Post(user, "invalid image format").Render(request.Context(), writer)
		return
	}

	filePath := fmt.Sprintf(FILE_PATH_FORMAT, imageID.String(), format)

	dst, err := os.Create(filePath)
	if err != nil {
		template.Post(user, "failed to write image to disk").Render(request.Context(), writer)
		return
	}
	defer dst.Close()

	if format == "png" {
		err = png.Encode(dst, image)
		if err != nil {
			template.Post(user, "failed to write image to disk").Render(request.Context(), writer)
			return
		}
	} else {
		err = jpeg.Encode(dst, image, nil)
		if err != nil {
			template.Post(user, "failed to write image to disk").Render(request.Context(), writer)
			return
		}
	}

	_, err = model.NewPost(user, imageID, fmt.Sprintf("/storage/%s.%s", imageID, format), request.FormValue("about"))
	if err != nil {
		log.Println(err)
		template.Post(user, "failed to create post in database").Render(request.Context(), writer)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}
