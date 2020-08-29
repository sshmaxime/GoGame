package websocket

import (
	"github.com/googollee/go-socket.io"
)

func wsResponse(socket socketio.Conn, data interface{}, path string) int {
	if err, ok := data.(error); ok {
		socket.Emit(ERROR, err.Error())
		return -1
	}
	socket.Emit(path, data)
	return 0
}
