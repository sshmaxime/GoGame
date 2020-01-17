package server

import (
	"fmt"
)

func (s *Server) GetNewGame(gameName string) (IGame, error) {
	gameCreatorFunction, found := s.GameCreatorFunctions[gameName]
	if !found {
		return nil, fmt.Errorf("game creator function for game [%v] not found", gameName)
	}

	newGame, err := LoadGameEngine(gameCreatorFunction)
	if err != nil {
		return nil, fmt.Errorf("impossible to load game engine for game [%v]: %v", gameName, err)
	}

	return newGame, nil
}
