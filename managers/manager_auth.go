package managers

import (
	"fmt"
	"github.com/GoGame/models"
)

type AuthManager struct {
	Users map[string]*models.User
}

func (m *AuthManager) Init() {
	m.Users = make(map[string]*models.User)

	_ = m.RegisterUser("admin", "admin")
}

func (m *AuthManager) RegisterUser(userID string, password string) error {
	if m.getUser(userID) != nil {
		return fmt.Errorf("UserID [%s] is already taken", userID)
	}
	newUser := models.User{UserID: userID, Password: password}
	m.Users[userID] = &newUser
	return nil
}

func (m *AuthManager) AuthenticateUser(userID string, password string) error {
	user := m.getUser(userID)
	if user == nil {
		return fmt.Errorf("UserID [%s] is not registered", userID)
	}
	if user.Password != password {
		return fmt.Errorf("Invalid password")
	}
	return nil
}

// Secret functions
func (m *AuthManager) getUser(userID string) *models.User {
	user, found := m.Users[userID]
	if !found {
		return nil
	}
	return user
}
