package server

import (
	"github.com/GoGame/network"
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestLoginParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	user, err := s.AuthManager.AuthenticateUser(req.UserID, req.Password)
	if err != nil {
		sendError(w, r, err)
		return
	}

	sendSuccessJSON(w, user)
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestRegisterParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	user := s.AuthManager.RegisterUser(req.UserID, req.Password)
	if user != nil {
		sendError(w, r, err)
		return
	}

	w.WriteHeader(200)
}
