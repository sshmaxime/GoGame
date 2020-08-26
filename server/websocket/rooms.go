package websocket

import (
	"github.com/GoGame/hub"
	"github.com/googollee/go-socket.io"
)

func initRooms(wsHandler *socketio.Server) {
	wsHandler.OnEvent("/", CREATE_ROOM_REQUEST, func(socket socketio.Conn, request CreateRoomRequest) int {
		return wsResponse(socket, hub.CreateRoomRequest(hub.GetClient(socket), request.Name), CREATE_ROOM_RESPONSE)
	})

	wsHandler.OnEvent("/", JOIN_ROOM_REQUEST, func(socket socketio.Conn, request JoinRoomRequest) int {
		return wsResponse(socket, hub.JoinRoomRequest(hub.GetClient(socket), request.Name), JOIN_ROOM_RESPONSE)
	})

	wsHandler.OnEvent("/", MESSAGE_ROOM_REQUEST, func(socket socketio.Conn, request MessageRoomRequest) int {
		return wsResponse(socket, hub.SendMessageToRoomRequest(hub.GetClient(socket), "/rooms", request.RoomName, request.Msg), MESSAGE_ROOM_RESPONSE)
	})
}
