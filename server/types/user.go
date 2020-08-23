package types

import (
	"github.com/googollee/go-socket.io"
)

type User struct {
	Username string `json:"username"`
}

type Client struct {
	Socket socketio.Conn
	Token  string
	User   *User
}
