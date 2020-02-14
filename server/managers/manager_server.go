package managers

import (
	"fmt"
	"github.com/GoGame/models"
	"math/rand"
	"strconv"
)

type ServerManager struct {
	gamesRoom map[string]map[string]*models.GameRoom
}

func ServerManagerConstructor() *ServerManager {
	this := &ServerManager{}
	this.gamesRoom = make(map[string]map[string]*models.GameRoom)
	return this
}

func (m *ServerManager) CreateGameRoom(gameName string, game models.IGame, user *models.User) (*models.GameRoom, error) {
	roomID := strconv.Itoa(rand.Int())

	newGameRoom := models.GameRoomConstructor(roomID, game, gameName)
	if err := newGameRoom.Join(user); err != nil {
		return nil, err
	}

	_, ok := m.gamesRoom[gameName]
	if !ok {
		m.gamesRoom[gameName] = make(map[string]*models.GameRoom)
	}

	m.gamesRoom[gameName][roomID] = newGameRoom
	return newGameRoom, nil
}

func (m *ServerManager) GetGameRoom(gameName string, gameRoomID string) (*models.GameRoom, error) {
	gameRoom, ok := m.gamesRoom[gameName][gameRoomID]
	if !ok {
		return nil, fmt.Errorf("gameRoomID [%v] from gameName [%v] doesn't exist", gameRoomID, gameName)
	}
	return gameRoom, nil
}
