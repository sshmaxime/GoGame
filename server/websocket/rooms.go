package websocket

import (
	"errors"
	"fmt"
	"github.com/googollee/go-socket.io"
)

func initRooms(wsHandler *socketio.Server) {
	wsHandler.OnConnect("/rooms", func(socket socketio.Conn) error {
		if _, ok := hub.clients[socket.ID()]; ok == false {
			return wsErrorAnDisconnect(socket, errors.New(""))
		}
		return nil
	})

	wsHandler.OnEvent("/rooms", "message", func(socket socketio.Conn, obj interface{}) error {
		fmt.Println(obj)
		return nil
	})
}
