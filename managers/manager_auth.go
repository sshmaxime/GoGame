package managers

import "github.com/GoGame/models"

type AuthManager struct {
	Users map[string]*models.User
}

func (m *AuthManager) Init() {
	m.Users = make(map[string]*models.User)

	m.AddUser("admin")
}

func (m *AuthManager) AddUser(userID string) {
	newUser := models.User{ID: userID}
	m.Users[userID] = &newUser
}

func (m *AuthManager) GetUser(userID string) *models.User {
	user, found := m.Users[userID]
	if !found {
		return nil
	}
	return user
}
