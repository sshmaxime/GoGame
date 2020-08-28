package hub

import (
	"github.com/googollee/go-socket.io"
)

func getRoom(name string) *Room {
	return rooms[name]
}

func GetClient(socket socketio.Conn) *Client {
	return clients[socket.ID()]
}
