package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		body, err := json.Marshal(h.GetStat())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
