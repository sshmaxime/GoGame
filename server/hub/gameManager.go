package hub

import (
	"errors"
	"fmt"
	"github.com/GoGame/types"
)

func CreateGameRequest(cli *Client, roomName string, gameName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return fmt.Errorf("room %v doesn't exist", roomName)
	}

	gameDefinition, err := getGameDefinition(gameName)
	if err != nil {
		return fmt.Errorf("game %v doesn't exist", gameName)
	}

	game := gameDefinition.CreatorFunction().(types.IGame)
	newGame := CreateGame(cli, game, gameDefinition)
	if err := room.addGame(newGame); err != nil {
		return err
	}

	msg := Message{
		From: "SERVER",
		Msg:  cli.User.Username + " created a game, join him !",
	}

	handler.BroadcastToRoom("/", room.Room.Name, MESSAGE_ROOM, msg)

	return JoinGameRequest(cli, roomName)
}

func JoinGameRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}
	room := getRoom(roomName)
	if room == nil {
		return fmt.Errorf("room %v doesn't exist", roomName)
	}
	if room.Game == nil {
		return fmt.Errorf("room %v doesn't have game", roomName)
	}

	if err := room.Game.addClient(cli); err != nil {
		return err
	}

	if len(room.Game.Clients) == room.Game.Game.GameDefinition.MaxPlayer {
		var players []string
		for _, player := range room.Game.Game.Users {
			players = append(players, player.Username)
		}
		room.Game.Game.Game.Init([]byte(""), players)

		// TODO send to players only
		handler.BroadcastToRoom("/", room.Room.Name, "GAME_STATE", room.Game.Game.Game.GetState())
	}

	return room.Game.Game.GameDefinition.ID
}

func LeaveGameRequest(cli *Client, roomName string) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return fmt.Errorf("room %v doesn't exist", roomName)
	}

	if room.Game == nil {
		return fmt.Errorf("room %v doesn't have game", roomName)
	}

	if err := room.Game.removeClient(cli); err != nil {
		return err
	}
	return room.Game.Game.GameDefinition.ID
}

func PlayGameRequest(cli *Client, roomName string, data []byte) interface{} {
	if cli == nil {
		return errors.New("client is not auth")
	}

	room := getRoom(roomName)
	if room == nil {
		return fmt.Errorf("room %v doesn't exist", roomName)
	}

	if room.Game == nil {
		return fmt.Errorf("room %v doesn't have game", roomName)
	}

	resp, err := room.Game.Game.Game.Play(data, cli.User.Username)
	if err != nil {
		return err
	}

	handler.BroadcastToRoom("/", room.Room.Name, "GAME_STATE", resp)

	return resp
}
