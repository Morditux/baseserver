package server

import (
	"io"
	"net/http"
	"sync"

	"htmx/sessions"
	"htmx/templates"
)

// ServerInstance represents the singleton instance of the server.
var ServerInstance *Server

// Server represents the HTTP server.
type Server struct {
	httpServer *http.Server
	router     *http.ServeMux
	templates  *templates.Templates
	data       map[string]interface{}
	sessions   *sessions.Sessions
	mutex      *sync.Mutex
	state      *State
}

// NewServer creates a new instance of the server.
func NewServer() *Server {
	templates := templates.NewTemplates()
	templates.AddSource("html")
	templates.AddSource("html/fragments")
	err := templates.Parse()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	ServerInstance = &Server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		router:    mux,
		templates: templates,
		mutex:     &sync.Mutex{},
		data:      make(map[string]interface{}),
		state: &State{
			Version: "0.1.0",
			mutex:   &sync.Mutex{},
		},
		sessions: sessions.NewSessions(),
	}
	return ServerInstance
}

// Start starts the server.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Stop stops the server.
func (s *Server) Stop() error {
	return s.httpServer.Close()
}

// HandleFunc registers a function to handle HTTP requests with the given pattern.
func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, handler)
}

// Handle registers a handler to handle HTTP requests with the given pattern.
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.router.Handle(pattern, handler)
}

// Templates returns the templates used by the server.
func (s *Server) Templates() *templates.Templates {
	return s.templates
}

// Data returns a copy of the server's data.
func (s *Server) Data() map[string]interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tmpData := make(map[string]interface{})
	for k, v := range s.data {
		tmpData[k] = v
	}
	return tmpData
}

// SetData sets the value of a key in the server's data.
func (s *Server) SetData(key string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

// GetData returns the value associated with the given key from the server's data.
func (s *Server) GetData(key string) interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.data[key]
}

// RemoveData removes the value associated with the given key from the server's data.
func (s *Server) RemoveData(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}

// State returns the server's state.
func (s *Server) State() *State {
	return s.state
}

// Sessions returns the sessions manager used by the server.
func (s *Server) Sessions() *sessions.Sessions {
	return s.sessions
}

// Render renders the specified template with the given data and writes the result to the response writer.
func (s *Server) Render(w io.Writer, template string, data map[string]interface{}) {
	s.templates.Execute(w, template, data)
}
