package server

type AuthManager struct {
	Users map[string]*User
}

func (a *AuthManager) Init() {
	a.Users = make(map[string]*User)
}

func (a *AuthManager) AddUser(userID string) {
	newUser := User{ID: userID}
	a.Users[userID] = &newUser
}

func (a *AuthManager) GetUser(userID string) *User {
	user, found := a.Users[userID]
	if !found {
		return nil
	}
	return user
}
