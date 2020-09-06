package websocket

import (
	"github.com/GoGame/hub"
	"github.com/googollee/go-socket.io"
)

func wsResponse(socket socketio.Conn, data interface{}, path string) int {
	if err, ok := data.(error); ok {
		socket.Emit(ERROR, err.Error())
		return -1
	}
	socket.Emit(path, data)

	if path == CREATE_ROOM_SUCCESS ||
		path == JOIN_ROOM_SUCCESS ||
		path == LEAVE_ROOM_SUCCESS ||
		path == LOGIN_SUCCESS {
		hub.SendState()
	}
	return 0
}
