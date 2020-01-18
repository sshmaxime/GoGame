package managers

import (
	"fmt"
	"github.com/GoGame/models"
)

type GameManager struct {
	gameCreatorFunctions map[string]func() interface{}
}

func (m *GameManager) Init() {
	m.gameCreatorFunctions = make(map[string]func() interface{})
}

func (m *GameManager) AddGame(gameName string, gameCreatorFunctions func() interface{}) {
	m.gameCreatorFunctions[gameName] = gameCreatorFunctions
}

func (m *GameManager) GetGame(gameName string) (models.IGame, error) {
	gameCreatorFunction, ok := m.gameCreatorFunctions[gameName]
	if !ok {
		return nil, fmt.Errorf("game [%v] not found", gameName)
	}

	gameAsInterface := gameCreatorFunction()

	game, ok := gameAsInterface.(models.IGame)
	if !ok {
		return nil, fmt.Errorf("game [%v] is from a wrong format", gameName)
	}

	return game, nil
}
