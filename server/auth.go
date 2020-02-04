package server

import (
	"github.com/GoGame/models"
)

func (s *Server) Authenticate(token string) *models.User {
	return s.AuthManager.AuthenticateWithToken(token)
}
