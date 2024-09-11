package server

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const (
	UPLOAD_LIMIT = 10 << 20 // 10 MB
	STORAGE_PATH = "./web/storage"
)

var (
	ALLOWED_TYPES = []string{
		"png",
		"jpg",
		"jpeg",
	}
)

func GetImageFromRequest(request *http.Request, name string) (image.Image, string, error) {
	file, _, err := request.FormFile(name)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	image, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	var found bool = false
	for _, allowedFormat := range ALLOWED_TYPES {
		if format == allowedFormat {
			found = true
		}
	}

	if !found {
		return nil, "", fmt.Errorf("image format not allowed")
	}

	return image, format, nil
}

func UploadImage(request *http.Request, name string) (uuid.UUID, string, error) {
	_, err := os.Stat(STORAGE_PATH)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(STORAGE_PATH, 0755)
			if err != nil {
				log.Println(err)
			}
		}
	}

	imageID, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, "", err
	}

	image, imageFormat, err := GetImageFromRequest(request, name)
	if err != nil {
		return uuid.Nil, "", err
	}

	dst, err := os.Create(fmt.Sprintf("%s/%s.%s", STORAGE_PATH, imageID.String(), imageFormat))
	if err != nil {
		return uuid.Nil, "", err
	}
	defer dst.Close()

	if imageFormat == "png" {
		err = png.Encode(dst, image)
		if err != nil {
			return uuid.Nil, "", err
		}
	} else {
		err = jpeg.Encode(dst, image, nil)
		if err != nil {
			return uuid.Nil, "", err
		}
	}

	return imageID, fmt.Sprintf("/storage/%s.%s", imageID, imageFormat), nil
}
