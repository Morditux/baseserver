package main

import (
	"runtime"

	"htmx/handlers"
	"htmx/server"
)

func main() {
	println(runtime.Version())
	server := server.NewServer()
	server.Handle("GET /pages/{page}", handlers.NewPageHandler())
	server.Handle("GET /css/{css}", handlers.NewCssHandler())
	server.Handle("POST /api/user/submit", handlers.NewUserHandler())
	server.Handle("/", handlers.NewIndexHandler())
	server.HandleFunc("GET /api/username", handlers.UserName)
	server.HandleFunc("GET /api/useremail", handlers.UserEmail)
	server.Start()
}
