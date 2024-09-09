package view

import (
	"log"
	"net/http"

	"github.com/maatko/society/api/model"
	authentication "github.com/maatko/society/internal/auth"
	"github.com/maatko/society/web/template/auth"
)

func POST_Register(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		auth.Register("internal server error", "").Render(request.Context(), writer)
		return
	}

	name := request.FormValue("name")
	password := request.FormValue("password")
	code := request.FormValue("code")

	invite, err := model.GetInviteByCode(code)
	if err != nil {
		auth.Register("invalid invite code!", "").Render(request.Context(), writer)
		return
	}

	if invite.UsedBy != nil {
		auth.Register("invite code already in use!", "").Render(request.Context(), writer)
		return
	}

	_, err = model.GetUserByName(name)
	if err == nil {
		log.Print(err)
		auth.Register("user already exists!", "").Render(request.Context(), writer)
		return
	}

	user, err := model.NewUser(name, password)
	if err != nil {
		auth.Register("internal server error!", "").Render(request.Context(), writer)
		return
	}

	invite.UsedBy = user
	err = invite.Update()
	if err != nil {
		user.Delete()
		auth.Register("internal server error!", "").Render(request.Context(), writer)
		return
	}

	err = authentication.Login(writer, user)
	if err != nil {
		auth.Register("failed to authenticate", "").Render(request.Context(), writer)
		return
	}

	http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
}

func GET_Register(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	auth.Register("", values.Get("code")).Render(request.Context(), writer)
}
