package websocket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
)

type Hub struct {
}

var hub *Hub

func Init() (wsHandler *socketio.Server, err error) {
	if wsHandler, err = socketio.NewServer(nil); err != nil {
		return nil, err
	}

	wsHandler.OnConnect("/", func(socket socketio.Conn) error {
		fmt.Println("connected:", socket.ID())
		return nil
	})

	wsHandler.OnDisconnect("/", func(socket socketio.Conn, str string) {
		fmt.Println("disconnected:", socket.ID())
	})

	return wsHandler, nil
}
