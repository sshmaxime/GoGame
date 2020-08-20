package main

import (
	"encoding/json"
)

type UpdateRequest struct {
	Y uint8 `json:"y"`
	X uint8 `json:"x"`
}

type State struct {
	Board [3][3]uint8 `json:"board"`
}

type Game struct {
	state State
}

func (g *Game) Init() {
	g.state.Board = [3][3]uint8{
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
		return InvalidRequest()
	}

	if g.state.Board[updateRequest.Y][updateRequest.X] != 0 {
		return InvalidRequest()
	}

	g.state.Board[updateRequest.Y][updateRequest.X] = playerID
	return nil
}

func (g *Game) GetState() interface{} {
	return State{
		Board: g.state.Board,
	}
}

func CreateGame() interface{} {
	return new(Game)
}
