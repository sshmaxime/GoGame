package managers

import (
	"github.com/GoGame/models"
)

type AuthManager struct {
	Users map[string]*models.User
}

func (m *AuthManager) Init() {
	m.Users = make(map[string]*models.User)

	m.Users["0000"] = &models.User{
		UserID:   "user",
		Password: "user",
	}
}

func (m *AuthManager) AuthenticateWithToken(token string) *models.User {
	user := m.getUserFromToken(token)
	if user == nil {
		return nil
	}
	return user
}

// Secret functions
func (m *AuthManager) getUserFromToken(token string) *models.User {
	user, found := m.Users[token]
	if !found {
		return nil
	}
	return user
}
