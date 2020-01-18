package models

import "sync"

type GameRoom struct {
	// Game data
	Game     IGame
	GameName string

	Mutex sync.Mutex
}
