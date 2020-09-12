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
	fmt.Println("done create game")
	return gameDefinition.ID
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
		room.Game.Game.Game.Init([]byte(""), []string{"player1"})
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

	resp, err := room.Game.Game.Game.Play(data, 1)
	if err != nil {
		return err
	}
	return resp
}
