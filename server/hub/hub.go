package hub

import (
	"github.com/GoGame/types"
	"github.com/googollee/go-socket.io"
)

var handler *socketio.Server
var clients map[string]*Client
var rooms map[string]*Room

func Init() (*socketio.Server, error) {
	var err error

	if handler, err = socketio.NewServer(nil); err != nil {
		return nil, err
	}

	clients = make(map[string]*Client)
	rooms = make(map[string]*Room)

	rooms["demo"] = &Room{
		Room: types.Room{
			Name:  "demo",
			Users: []*types.User{},
		},
		Clients: map[string]*Client{},
	}
	return handler, nil
}

func AddClient(cli *Client) error {
	clients[cli.Socket.ID()] = cli
	return nil
}

func IsClientAuth(socket socketio.Conn) bool {
	_, ok := clients[socket.ID()]
	return ok
}
