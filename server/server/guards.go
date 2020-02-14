package server

import (
	"fmt"
	"github.com/GoGame/models"
	"net/http"
)

func (s *Server) authGuard(r *http.Request) (*models.User, error) {
	var user *models.User
	if user = s.AuthManager.AuthenticateWithToken(r.Header.Get("X-Token")); user == nil {
		return nil, fmt.Errorf("invalid token")
	}
	return user, nil
}

// Game Guards
func (s *Server) GameGuard(r *http.Request, request interface{}) (*models.User, error) {
	var user *models.User
	var err error

	if user, err = s.authGuard(r); err != nil {
		return nil, err
	}
	if err = ReadBody(r.Body, &request); err != nil {
		return nil, fmt.Errorf("invalid request")
	}
	return user, nil
}
