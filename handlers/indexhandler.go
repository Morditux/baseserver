package handlers

import (
	"log/slog"
	"net/http"

	"htmx/server"
)

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("IndexHandler.ServeHTTP")
	session := server.ServerInstance.Sessions().GetSession(w, r)
	session.Data["Connected"] = true
	data := MakeData()
	data["Title"] = "Htmx test"
	data["Message"] = "Hello, World!"

	server.ServerInstance.Templates().Execute(w, "index", data)
	server.ServerInstance.SetData("Message", "This is a message from the server")
}
