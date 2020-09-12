package main

import (
	"encoding/json"
)

type Action struct {
	Y uint `json:"y"`
	X uint `json:"x"`
}

type Game struct {
	Board     [3][3]uint `json:"board"`
	WhoToPlay []uint     `json:"who_to_play"`
	Victory   uint       `json:"victory"`

	Players map[string]uint `json:"players"`
}

// External
func CreateGame() interface{} {
	return new(Game)
}

// Interface
func (g *Game) Init(_ []byte, players []string) {
	g.Board = [3][3]uint{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	g.Victory = 0
	g.Players = make(map[string]uint)
	for index, player := range players {
		index++ // First player start at 1
		g.Players[player] = uint(index)
		g.WhoToPlay = append(g.WhoToPlay, uint(index))
	}
}

func (g *Game) Play(actionAsBytes []byte, playerID string) (interface{}, error) {
	var action Action
	err := json.Unmarshal(actionAsBytes, &action)
	if err != nil {
		return Game{}, InvalidRequest("impossible to parse action")
	}
	return g.processAction(&action, playerID)
}

func (g *Game) GetState() interface{} { return *g }

//

// ProcessAction
func (g *Game) processAction(action *Action, playerID string) (Game, error) {
	// Check errors
	if err := g.isActionAllowed(action, playerID); err != nil {
		return Game{}, InvalidRequest(err.Error())
	}

	// Process action
	g.Board[action.Y][action.X] = g.Players[playerID]
	// Put the current player at the end of the waiting to play list
	g.WhoToPlay = append(g.WhoToPlay[1:], g.WhoToPlay[0])

	// Tick
	return g.tick(), nil
}

//

// Tick
func (g *Game) tick() Game {
	// TODO Check end
	return *g
}

//

// Utils
func (g *Game) isActionAllowed(action *Action, playerID string) error {
	if g.Victory != 0 {
		return InvalidRequest("game is over")
	}
	if g.WhoToPlay[0] != g.Players[playerID] {
		return InvalidRequest("not your turn to play")
	}
	if action.Y > 3 || action.Y < 0 ||
		action.X > 3 || action.X < 0 {
		return InvalidRequest("action out of range")
	}
	if g.Board[action.Y][action.X] != 0 {
		return InvalidRequest("position already filled")
	}
	return nil
}

//
