package server

import (
	"net/http"
)

type Handler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

// Default routes
func (s *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
