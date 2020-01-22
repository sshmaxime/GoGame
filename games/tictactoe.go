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

func (g *Game) Update(updateRequestAsBytes []byte, userID uint8) {
	var updateRequest UpdateRequest
	err := json.Unmarshal(updateRequestAsBytes, &updateRequest)
	if err != nil {
		LogUpdateError(err)
		return
	}

	if updateRequest.Y > 3 || updateRequest.X > 3 {
		LogUpdateError(errors.New("invalid parameters"))
		return
	}

	if g.state.Board[updateRequest.Y][updateRequest.X] != 0 {
		LogUpdateError(errors.New("invalid position"))
		return
	}

	g.state.Board[updateRequest.Y][updateRequest.X] = userID
	LogUpdateInfo("Someone played")
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
