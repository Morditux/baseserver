package server

import "sync"

type State struct {
	Version   string
	lastError string
	mutex     *sync.Mutex
}

func (s *State) SetError(err string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.lastError = err
}

func (s *State) LastError() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.lastError
}
