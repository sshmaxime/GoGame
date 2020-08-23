package database

import (
	"fmt"
	"github.com/GoGame/config"
	"github.com/GoGame/types"
	"github.com/GoGame/utils"
)

type GameDatabase struct {
	games map[string]*types.GameDefinition
}

var gameDatabase GameDatabase

func initGameDatabase() (err error) {
	gameDatabase.games = map[string]*types.GameDefinition{}

	for _, game := range config.GetGames() {
		// Load the game engine creator function
		if game.CreatorFunction, err = utils.LoadGameEngineCreatorFunction(game.LibPath); err != nil {
			return err
		}
		gameDatabase.games[game.ID] = &game
	}
	return nil
}

func GetAllGames() (map[string]*types.GameDefinition, error) {
	return gameDatabase.games, nil
}

func GetGameByID(ID string) (game *types.GameDefinition, err error) {
	errorMsg := "error while getting game"

	if game = gameDatabase.games[ID]; game == nil {
		return nil, fmt.Errorf("%v: game [%v] doesn't exist", errorMsg, ID)
	}
	return game, nil
}
