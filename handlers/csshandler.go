package handlers

import (
	"io"
	"net/http"
	"os"
)

type CssHandler struct {
}

func NewCssHandler() *CssHandler {
	return &CssHandler{}
}

func (h *CssHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	css := r.PathValue("css")
	fileName := "css/" + css
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
