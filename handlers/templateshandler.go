package handlers

import (
	"net/http"

	"htmx/server"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page := r.PathValue("page")
	if page == "" {
		page = "notfound"
	}
	server.ServerInstance.Templates().Execute(w, page, server.ServerInstance.Data())
}
