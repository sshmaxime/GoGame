package managers

import (
	"fmt"
	"github.com/GoGame/models"
	"time"
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
		return fmt.Errorf("userID [%s] is already taken", userID)
	}
	newUser := models.User{UserID: userID, Password: password}
	m.Users[userID] = &newUser
	return nil
}

func (m *AuthManager) AuthenticateUser(userID string, password string) (models.User, error) {
	user := m.getUser(userID)
	if user == nil {
		return models.User{}, fmt.Errorf("userID [%s] is not registered", userID)
	}
	if user.Password != password {
		return models.User{}, fmt.Errorf("invalid password")
	}
	user.Token = userID + string(time.Now().Nanosecond())
	return *user, nil
}

func (m *AuthManager) AuthenticateWithToken(userID string, token string) error {
	user := m.getUser(userID)
	if user == nil {
		return fmt.Errorf("userID [%s] is not registered", userID)
	}
	if user.Token != token {
		return fmt.Errorf("invalid token")
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
