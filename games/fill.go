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
	Board [][]string `json:"board"`
}

type Game struct {
	state State
	users map[string]uint8
}

func (g *Game) Init() {
	g.state.Board = [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}
}

func (g *Game) Play(updateRequestAsBytes []byte, userID string) error {
	var updateRequest UpdateRequest
	err := json.Unmarshal(updateRequestAsBytes, &updateRequest)
	if err != nil {
		return LogUpdateError(err)
	}

	if updateRequest.Y > 3 || updateRequest.X > 3 {
		return LogUpdateError(errors.New("invalid parameters"))
	}

	if g.state.Board[updateRequest.Y][updateRequest.X] != "" {
		return LogUpdateError(errors.New("invalid position"))
	}

	g.state.Board[updateRequest.Y][updateRequest.X] = userID
	LogUpdateInfo("Someone played")
	return nil
}

func (g *Game) GetState() interface{} {
	state := State{
		Board: g.state.Board,
	}
	return state
}

func CreateGame() interface{} {
	return new(Game)
}
