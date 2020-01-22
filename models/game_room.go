package models

import (
	"fmt"
	"sync"
)

type GameRoom struct {
	ID string

	// Game data
	Game     IGame
	GameName string

	// User
	Users map[string]*User

	Mutex sync.Mutex
}

func (room *GameRoom) Init() {
	room.Users = make(map[string]*User)
}

func (room *GameRoom) AddUser(newUser *User) error {
	_, found := room.Users[newUser.UserID]
	if found {
		return fmt.Errorf("user [%s] is already in the game room", newUser.UserID)
	}
	room.Users[newUser.UserID] = newUser
	return nil
}

func (room *GameRoom) RemoveUser(userID string) error {
	_, found := room.Users[userID]
	if !found {
		return fmt.Errorf("user [%s] is not in the game room", userID)
	}
	delete(room.Users, userID)
	return nil
}
