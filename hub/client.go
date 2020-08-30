package hub

import (
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

type Client struct {
	Socket socketio.Conn
	User   *types.User
	Rooms  map[string]bool
}

func CreateClient(socket socketio.Conn, user *types.User) *Client {
	return &Client{
		Socket: socket,
		User:   user,
		Rooms:  make(map[string]bool),
	}
}

func (cli *Client) joinRoom(roomName string) error {
	cli.Rooms[roomName] = true
	cli.Socket.Join(roomName)
	return nil
}

func (cli *Client) leaveRoom(roomName string) error {
	delete(cli.Rooms, roomName)
	cli.Socket.Leave(roomName)
	return nil
}
