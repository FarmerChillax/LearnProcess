package server

import (
	"http-cache-service-demo/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":233", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
