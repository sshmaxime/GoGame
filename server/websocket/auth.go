package websocket

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

func initAuth(wsHandler *socketio.Server) {
	wsHandler.OnEvent("/", types.LOGIN_REQUEST, func(socket socketio.Conn, request types.LoginRequest) error {
		var user *types.User
		var err error

		if user, err = database.AuthenticateUser(request.Username, request.Password); err != nil {
			return wsError(socket, err)
		}

		client := types.Client{
			Socket: socket,
			User:   user,
		}

		hub.clients[socket.ID()] = client
		socket.Emit(types.LOGIN_RESPONSE, client.User)
		return nil
	})
}
