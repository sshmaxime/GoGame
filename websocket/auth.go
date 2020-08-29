package websocket

import (
	"github.com/GoGame/database"
	"github.com/GoGame/hub"
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

func initAuth(wsHandler *socketio.Server) {
	wsHandler.OnEvent("/", LOGIN_REQUEST, func(socket socketio.Conn, request LoginRequest) int {
		var user *types.User
		var err error

		if user, err = database.AuthenticateUser(request.Username, request.Password); err != nil {
			return wsResponse(socket, err, "")
		}

		newClient := hub.CreateClient(socket, user)
		if err = hub.AddClient(newClient); err != nil {
			return wsResponse(socket, err, "")
		}

		return wsResponse(socket, newClient.User, LOGIN_SUCCESS)
	})

	wsHandler.OnEvent("/", REGISTER_REQUEST, func(socket socketio.Conn, request LoginRequest) int {
		var user *types.User
		var err error

		if user, err = database.CreateUser(request.Username, request.Password); err != nil {
			return wsResponse(socket, err, "")
		}

		newClient := hub.CreateClient(socket, user)
		if err = hub.AddClient(newClient); err != nil {
			return wsResponse(socket, err, "")
		}

		return wsResponse(socket, newClient.User, REGISTER_SUCCESS)
	})
}
