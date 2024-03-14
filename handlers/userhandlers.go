package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"htmx/model"
	"htmx/server"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	valid := true
	session := server.ServerInstance.Sessions().GetSession(w, r)
	r.ParseForm()
	form := r.PostForm
	username := form.Get("name")
	email := form.Get("email")
	data := MakeData()

	if username == "" {
		data["InvalidName"] = true
		valid = false
	}
	if email == "" {
		data["InvalidEmail"] = true
		valid = false
	}
	if !valid {
		server.ServerInstance.Render(w, "form", data)
		return
	}

	user := model.User{
		Name:  username,
		Email: email,
	}
	session.Data["user"] = user
	SetTrigger(w, "userConnected")
	server.ServerInstance.Render(w, "body", data)
}

func UserName(w http.ResponseWriter, r *http.Request) {
	session := server.ServerInstance.Sessions().GetSession(w, r)
	user, ok := session.Data["user"].(model.User)
	if !ok {
		slog.Info(r.RequestURI, "No user")
		fmt.Fprint(w, "No user")
		return
	}
	fmt.Fprintf(w, user.Name)
}

func UserEmail(w http.ResponseWriter, r *http.Request) {
	session := server.ServerInstance.Sessions().GetSession(w, r)
	user, ok := session.Data["user"].(model.User)
	if !ok {
		slog.Info(r.RequestURI, "No user")
		fmt.Fprint(w, "No user")
		return
	}
	fmt.Fprintf(w, user.Email)
}
