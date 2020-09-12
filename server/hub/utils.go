package hub

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

func getRoom(name string) *Room {
	return rooms[name]
}

func GetClient(socket socketio.Conn) *Client {
	return clients[socket.ID()]
}

func getGameDefinition(id string) (*types.GameDefinition, error) {
	return database.GetGameByID(id)
}
