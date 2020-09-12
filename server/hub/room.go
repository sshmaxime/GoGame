package hub

import (
	"fmt"
	"github.com/GoGame/types"
)

type Room struct {
	Room    *types.Room
	Game    *Game
	Clients map[string]*Client
}

func CreateRoom(cli *Client, roomName string) *Room {
	return &Room{
		Room: &types.Room{
			Name: roomName,
			Users: map[string]*types.User{
				cli.Socket.ID(): cli.User,
			},
		},
		Clients: map[string]*Client{
			cli.Socket.ID(): cli,
		},
	}
}

func (room *Room) addClient(cli *Client) error {
	room.Clients[cli.Socket.ID()] = cli
	room.Room.Users[cli.Socket.ID()] = cli.User
	return nil
}

func (room *Room) addGame(game *Game) error {
	room.Game = game
	return nil
}

func (room *Room) removeClient(cli *Client) error {
	cli, ok := room.Clients[cli.Socket.ID()]
	if !ok {
		return fmt.Errorf("%v is not in room %v", cli.User.Username, room.Room.Name)
	}
	delete(room.Clients, cli.Socket.ID())
	delete(room.Room.Users, cli.Socket.ID())
	return nil
}
