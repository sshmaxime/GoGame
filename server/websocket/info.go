package websocket

import (
	"github.com/googollee/go-socket.io"
)

func initInfo(wsHandler *socketio.Server) {
	wsHandler.OnConnect("/info", func(socket socketio.Conn) error {
		return nil
	})

	wsHandler.OnEvent("/info", "message", func(socket socketio.Conn, obj interface{}) error {
		return nil
	})
}
