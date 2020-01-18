package server

import (
	"github.com/GoGame/network"
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
	req, err := network.HandleRequestInitParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	user := s.AuthManager.GetUser(req.UserID)
	if user == nil {
		sendError(w, r, err)
		return
	}
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	_, err := network.HandleRequestUpdateParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}
}
func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
}

func sendError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), 500)
}
