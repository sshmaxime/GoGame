package hub

import (
	"errors"
	"fmt"
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

	// Default room
	rooms["general"] = &Room{
		Room:    &types.Room{
			Name:  "general",
			Users: make(map[string]*types.User),
		},
		Clients: make(map[string]*Client),
	}

	return handler, nil
}

func AddClient(cli *Client) error {
	for _, client := range clients {
		if client.User.Username == cli.User.Username {
			return fmt.Errorf("user %v already logged in", client.User.Username)
		}
	}
	clients[cli.Socket.ID()] = cli
	return nil
}

func RemoveClient(cli *Client) error {
	if cli == nil {
		return errors.New("client is not auth")
	}

	for roomName, _ := range cli.Rooms {
		for _, room := range rooms {
			if roomName == room.Room.Name {
				cli.leaveRoom(roomName)
				room.removeClient(cli)
			}
		}
	}
	delete(clients, cli.Socket.ID())
	return nil
}
