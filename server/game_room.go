package server

import "sync"

type GameRoom struct {
	// Game data
	GameEngine IGame

	Mutex sync.Mutex
}
