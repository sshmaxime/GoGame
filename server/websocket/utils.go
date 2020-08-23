package websocket

import (
	"fmt"
	"github.com/googollee/go-socket.io"
)

func wsError(socket socketio.Conn, err error) error {
	socket.Emit("error", err.Error())
	return err
}

func wsErrorAnDisconnect(socket socketio.Conn, err error) error {
	socket.Emit("error", err.Error())
	socket.Close()
	fmt.Println("error spotted")
	return err
}
