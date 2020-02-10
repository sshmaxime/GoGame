package models

import (
	"fmt"
)

type GameRoom struct {
	ID string

	// Game data
	game     IGame
	GameName string

	// User
	users          map[string]*User
	players        map[string]uint8
	playersCounter uint8
}

func GameRoomConstructor(ID string, game IGame, gameName string) *GameRoom {
	this := &GameRoom{}
	this.ID = ID
	this.game = game
	this.GameName = gameName
	this.users = make(map[string]*User)
	this.players = make(map[string]uint8)
	this.playersCounter = 1
	return this
}

func (room *GameRoom) Join(newUser *User) error {
	_, found := room.users[newUser.UserID]
	if found {
		return fmt.Errorf("user [%s] is already in the game room", newUser.UserID)
	}
	room.users[newUser.UserID] = newUser
	room.players[newUser.UserID] = room.playersCounter
	room.playersCounter++
	return nil
}

func (room *GameRoom) Leave(userID string) error {
	_, found := room.users[userID]
	if !found {
		return fmt.Errorf("user [%s] is not in the game room", userID)
	}
	delete(room.users, userID)
	delete(room.players, userID)
	return nil
}

func (room *GameRoom) Play(data []byte, userID string) error {
	playerID, found := room.players[userID]
	if !found {
		return fmt.Errorf("user [%s] is not in the game room", userID)
	}
	if err := room.game.Play(data, playerID); err != nil {
		return fmt.Errorf("error while playing %v", userID)
	}
	return nil
}

func (room *GameRoom) GetState(userID string) (interface{}, error) {
	_, found := room.players[userID]
	if !found {
		return nil, fmt.Errorf("user [%s] is not in the game room", userID)
	}
	return room.game.GetState(), nil
}
