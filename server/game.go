package server

import (
	"errors"
	"plugin"
)

type IGame interface {
	Init()
	Get() int
	Add()
}

func LoadGameEngineCreatorFunction(gameEngineLibPath string) (func() interface{}, error) {
	plug, err := plugin.Open(gameEngineLibPath)
	if err != nil {
		return nil, errors.New("impossible to find plugin")
	}

	symbol, err := plug.Lookup("CreateGame")
	if err != nil {
		return nil, errors.New("impossible to find symbol")
	}

	creatorFunc, ok := symbol.(func() interface{})
	if !ok {
		return nil, errors.New("unexpected type from module symbol")
	}

	return creatorFunc, nil
}

func LoadGameEngine(creatorFunc func() interface{}) (IGame, error) {
	gameAsInterface := creatorFunc()
	game, ok := gameAsInterface.(IGame)
	if !ok {
		return nil, errors.New("unexpected type from create game")
	}
	return game, nil
}
