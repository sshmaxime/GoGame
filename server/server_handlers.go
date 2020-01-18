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

// Game routes
func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	_, err := handleRequestInitParsing(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	_, err := handleRequestUpdateParsing(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
}
