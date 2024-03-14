package sessions

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

// Sessions represents a collection of sessions.
type Sessions struct {
	s     map[string]Session
	mutex *sync.Mutex
}

// Session represents a session with associated data.
type Session struct {
	Data map[string]interface{}
}

// NewSessions creates a new Sessions instance.
func NewSessions() *Sessions {
	return &Sessions{
		s:     make(map[string]Session),
		mutex: &sync.Mutex{},
	}
}

// GetSession retrieves the session associated with the request.
// If the session already exists, it is returned.
// Otherwise, a new session is created and returned.
func (s *Sessions) GetSession(w http.ResponseWriter, r *http.Request) Session {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if the session already exists
	cookie, err := r.Cookie("htmxsession")
	if err == nil {
		session, ok := s.s[cookie.Value]
		if ok {
			return session
		}
		session = Session{
			Data: make(map[string]interface{}),
		}
		s.s[cookie.Value] = session
		return session
	}

	session := Session{
		Data: make(map[string]interface{}),
	}
	id := SetCookie(w, "htmxsession")
	s.s[id] = session
	return session
}

// SetCookie sets a cookie with the given name and returns its value.
func SetCookie(w http.ResponseWriter, name string) string {
	id, err := uuid.NewRandom()
	if err != nil {
		slog.Error(err.Error())
		return ""
	}
	cookie := http.Cookie{
		Name:     name,
		Value:    id.String(),
		Path:     "/", // This is important
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	return id.String()
}
