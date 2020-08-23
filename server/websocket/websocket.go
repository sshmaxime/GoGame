package websocket

import (
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

type Hub struct {
	clients map[string]types.Client
}

var hub *Hub

func Init() (wsHandler *socketio.Server, err error) {
	hub = &Hub{
		clients: make(map[string]types.Client),
	}

	if wsHandler, err = socketio.NewServer(nil); err != nil {
		return nil, err
	}

	wsHandler.OnConnect("/", func(socket socketio.Conn) error {
		return nil
	})
	wsHandler.OnDisconnect("/", func(socket socketio.Conn, str string) {
		delete(hub.clients, socket.ID())
	})

	initAuth(wsHandler)
	initInfo(wsHandler)
	initRooms(wsHandler)

	return wsHandler, nil
}
