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

func GameRoomConstructor(ID string, game IGame, gameName string) *GameRoom {
	this := &GameRoom{}
	this.ID = ID
	this.Game = game
	this.GameName = gameName
	this.Users = make(map[string]*User)
	return this
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
