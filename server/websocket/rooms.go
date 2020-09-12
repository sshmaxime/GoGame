package websocket

import (
	"fmt"
	"github.com/GoGame/hub"
	"github.com/googollee/go-socket.io"
)

func initRooms(wsHandler *socketio.Server) {
	// Room
	wsHandler.OnEvent("/", CREATE_ROOM_REQUEST, func(socket socketio.Conn, request CreateRoomRequest) int {
		return wsResponse(socket, hub.CreateRoomRequest(hub.GetClient(socket), request.Name), CREATE_ROOM_SUCCESS)
	})
	wsHandler.OnEvent("/", JOIN_ROOM_REQUEST, func(socket socketio.Conn, request JoinRoomRequest) int {
		return wsResponse(socket, hub.JoinRoomRequest(hub.GetClient(socket), request.Name), JOIN_ROOM_SUCCESS)
	})
	wsHandler.OnEvent("/", LEAVE_ROOM_REQUEST, func(socket socketio.Conn, request LeaveRoomRequest) int {
		return wsResponse(socket, hub.LeaveRoomRequest(hub.GetClient(socket), request.Name), LEAVE_ROOM_SUCCESS)
	})
	wsHandler.OnEvent("/", MESSAGE_ROOM_REQUEST, func(socket socketio.Conn, request MessageRoomRequest) int {
		return wsResponse(socket, hub.SendMessageToRoomRequest(hub.GetClient(socket), "/", request.RoomName, request.Msg), MESSAGE_ROOM_SUCCESS)
	})

	// Game
	wsHandler.OnEvent("/", CREATE_GAME_REQUEST, func(socket socketio.Conn, request CreateGameRequest) int {
		return wsResponse(socket, hub.CreateGameRequest(hub.GetClient(socket), request.RoomName, request.Name), CREATE_GAME_SUCCESS)
	})
	wsHandler.OnEvent("/", JOIN_GAME_REQUEST, func(socket socketio.Conn, request JoinGameRequest) int {
		return wsResponse(socket, hub.JoinGameRequest(hub.GetClient(socket), request.RoomName), JOIN_GAME_SUCCESS)
	})
	wsHandler.OnEvent("/", LEAVE_GAME_REQUEST, func(socket socketio.Conn, request LeaveGameRequest) int {
		return wsResponse(socket, hub.LeaveGameRequest(hub.GetClient(socket), request.RoomName), LEAVE_GAME_SUCCESS)
	})
	wsHandler.OnEvent("/", PLAY_GAME_REQUEST, func(socket socketio.Conn, request PlayGameRequest) int {
		fmt.Println("la")
		return wsResponse(socket, hub.PlayGameRequest(hub.GetClient(socket), request.RoomName, []byte(request.Data)), PLAY_GAME_SUCCESS)
	})
}
