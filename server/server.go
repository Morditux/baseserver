package server

import (
	"net/http"
	"sync"

	"htmx/sessions"
	"htmx/templates"
)

var ServerInstance *Server

type Server struct {
	httpServer *http.Server
	router     *http.ServeMux
	templates  *templates.Templates
	data       map[string]interface{}
	sessions   *sessions.Sessions
	mutex      *sync.Mutex
	state      *State
}

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

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}

func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, handler)
}

func (s *Server) Handle(pattern string, handler http.Handler) {
	s.router.Handle(pattern, handler)
}

func (s *Server) Templates() *templates.Templates {
	return s.templates
}

func (s *Server) Data() map[string]interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tmpData := make(map[string]interface{})
	for k, v := range s.data {
		tmpData[k] = v
	}
	return tmpData
}

func (s *Server) SetData(key string, value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[key] = value
}

func (s *Server) GetData(key string) interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.data[key]
}

func (s *Server) RemoveData(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}

func (s *Server) State() *State {
	return s.state
}

func (s *Server) Sessions() *sessions.Sessions {
	return s.sessions
}

func (s *Server) Render(w http.ResponseWriter, template string, data map[string]interface{}) {
	s.templates.Execute(w, template, data)
}
