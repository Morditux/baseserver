package handlers

import (
	"net/http"
	"strings"

	"htmx/server"
)

type BaseHandler struct {
	server *server.Server
}

func NewData() map[string]interface{} {
	return make(map[string]interface{})
}

func MakeData() map[string]interface{} {
	return make(map[string]interface{})
}

func SetTrigger(w http.ResponseWriter, triggers ...string) {
	if len(triggers) == 0 {
		return
	}

	if len(triggers) == 1 {
		w.Header().Add("HX-Trigger", triggers[0])
		return
	}

	var sb strings.Builder
	for i, trigger := range triggers {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(trigger)
	}
	w.Header().Add("HX-Trigger", sb.String())
}
