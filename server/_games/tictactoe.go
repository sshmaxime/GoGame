package main

import (
	"encoding/json"
)

type Action struct {
	Y uint8 `json:"y"`
	X uint8 `json:"x"`
}

type Game struct {
	Board [3][3]uint8 `json:"board"`
	WhoToPlay uint8 `json:"who_to_play"`
	Victory uint8 `json:"victory"`
}
// External
func CreateGame() interface{} {
	return new(Game)
}

// Interface
func (g *Game) Init(_ []byte, _ uint8) {
	g.Board = [3][3]uint8{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
}

func (g *Game) Play(actionAsBytes []byte, playerID uint8) (interface{}, error) {
	var action Action
	err := json.Unmarshal(actionAsBytes, &action)
	if err != nil {
		return Game{}, InvalidRequest()
	}
	return g.processAction(&action, playerID)
}

func (g *Game) GetState() interface{} {
	return *g
}
//

// ProcessAction
func (g *Game) processAction(action *Action, playerID uint8) (Game, error) {
	// Check errors
	if err := g.isActionAllowed(action); err != nil {
		return  Game{}, InvalidRequest()
	}

	// Process action
	g.Board[action.Y][action.X] = playerID

	// Tick
	return g.tick(), nil
}
//

// Tick
func (g *Game) tick() Game {
	// TODO
	return *g
}
//

// Utils
func (g *Game) isActionAllowed(action *Action) error {
	if action.Y > 3 || action.Y < 1 ||
		action.X > 3 || action.X < 1 {
		return InvalidRequest()
	}
	if g.Board[action.Y][action.X] != 0 {
		return InvalidRequest()
	}
	return nil
}
//