package hub

import (
	"fmt"
	"github.com/GoGame/types"
)

type Game struct {
	Game    *types.Game
	Clients map[string]*Client
}

func CreateGame(cli *Client, game types.IGame, gameDefinition *types.GameDefinition) *Game {
	return &Game{
		Game: &types.Game{
			Users: map[string]*types.User{
				cli.Socket.ID(): cli.User,
			},
			GameDefinition: gameDefinition,
			Game:           game,
		},
		Clients: map[string]*Client{
			cli.Socket.ID(): cli,
		},
	}
}

func (game *Game) addClient(cli *Client) error {
	game.Clients[cli.Socket.ID()] = cli
	game.Game.Users[cli.Socket.ID()] = cli.User
	return nil
}

func (game *Game) removeClient(cli *Client) error {
	cli, ok := game.Clients[cli.Socket.ID()]
	if !ok {
		return fmt.Errorf("%v is not in game", cli.User.Username)
	}
	delete(game.Clients, cli.Socket.ID())
	delete(game.Game.Users, cli.Socket.ID())
	return nil
}
