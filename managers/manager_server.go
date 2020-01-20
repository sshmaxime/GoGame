package managers

import (
	"fmt"
	"github.com/GoGame/models"
	"time"
)

type _gameName string
type _gameRoomID string

type ServerManager struct {
	gamesRoom map[_gameName]map[_gameRoomID]*models.GameRoom
}

func (m *ServerManager) Init() {
	m.gamesRoom = make(map[_gameName]map[_gameRoomID]*models.GameRoom)
}

func (m *ServerManager) CreateGameRoom(gameName string, game models.IGame) models.GameRoom {
	newGameRoom := new(models.GameRoom)

	newGameRoom.Game = game
	newGameRoom.GameName = gameName

	_, ok := m.gamesRoom[_gameName(gameName)]
	if !ok {
		m.gamesRoom[_gameName(gameName)] = make(map[_gameRoomID]*models.GameRoom)
	}

	roomID := string(time.Now().Nanosecond())
	m.gamesRoom[_gameName(gameName)][_gameRoomID(roomID)] = newGameRoom
	return *newGameRoom
}

func (m *ServerManager) GetGameRoom(gameName string, gameRoomID string) (*models.GameRoom, error) {
	gameRoom, ok := m.gamesRoom[_gameName(gameName)][_gameRoomID(gameRoomID)]
	if !ok {
		return nil, fmt.Errorf("gameRoomID [%v] from gameName [%v] doesn't exist", gameRoomID, gameName)
	}
	return gameRoom, nil
}
