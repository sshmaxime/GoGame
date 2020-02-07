package main

import (
	"encoding/json"
	"errors"
)

type UpdateRequest struct {
	Y uint8 `json:"y"`
	X uint8 `json:"x"`
}

type State struct {
	Board [][]uint8 `json:"board"`
}

type Game struct {
	state State
}

func (g *Game) Init() {
	g.state.Board = [][]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

func (g *Game) Play(updateRequestAsBytes []byte, playerID uint8) error {
	var updateRequest UpdateRequest
	err := json.Unmarshal(updateRequestAsBytes, &updateRequest)
	if err != nil {
		return err
	}

	if updateRequest.Y > 3 || updateRequest.X > 3 {
		return errors.New("invalid parameters")
	}

	if g.state.Board[updateRequest.Y][updateRequest.X] != 0 {
		return errors.New("invalid position")
	}

	g.state.Board[updateRequest.Y][updateRequest.X] = playerID
	return nil
}

func (g *Game) State() interface{} {
	state := State{
		Board: g.state.Board,
	}
	return state
}

func CreateGame() interface{} {
	return new(Game)
}
