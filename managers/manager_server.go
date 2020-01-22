package managers

import (
	"fmt"
	"github.com/GoGame/models"
	"math/rand"
	"strconv"
)

type _gameName string
type _gameRoomID string

type ServerManager struct {
	gamesRoom map[_gameName]map[_gameRoomID]*models.GameRoom
}

func (m *ServerManager) Init() {
	m.gamesRoom = make(map[_gameName]map[_gameRoomID]*models.GameRoom)
}

func (m *ServerManager) CreateGameRoom(gameName string, game models.IGame, user *models.User) models.GameRoom {
	newGameRoom := new(models.GameRoom)
	newGameRoom.Init()

	newGameRoom.Game = game
	newGameRoom.GameName = gameName
	_ = newGameRoom.AddUser(user)

	_, ok := m.gamesRoom[_gameName(gameName)]
	if !ok {
		m.gamesRoom[_gameName(gameName)] = make(map[_gameRoomID]*models.GameRoom)
	}

	roomID := strconv.Itoa(rand.Int())
	newGameRoom.ID = roomID
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
