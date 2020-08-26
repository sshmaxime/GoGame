package hub

import (
	"github.com/GoGame/types"
)

type Room struct {
	Room    types.Room
	Clients map[string]*Client
}

func CreateRoom(cli *Client, roomName string) *Room {
	return &Room{
		Room: types.Room{
			Name: roomName,
			Users: []*types.User{
				cli.User,
			},
		},
		Clients: map[string]*Client{
			cli.Socket.ID(): cli,
		},
	}
}

func (room *Room) addClient(cli *Client) error {
	room.Clients[cli.Socket.ID()] = cli
	room.Room.Users = append(room.Room.Users, cli.User)
	return nil
}
