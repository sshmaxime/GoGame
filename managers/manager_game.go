package managers

import (
	"github.com/GoGame/models"
)

type GameManager struct {
	gameCreatorFunctions map[string]func() interface{}
	gamesRoom            map[string]map[string]*models.GameRoom
}

func (m *GameManager) Init() {
	m.gameCreatorFunctions = make(map[string]func() interface{})
	m.gamesRoom = make(map[string]map[string]*models.GameRoom)
}

func (m *GameManager) AddGame(gameName string, gameCreatorFunctions func() interface{}) {
	m.gameCreatorFunctions[gameName] = gameCreatorFunctions
}
