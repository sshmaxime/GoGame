package main

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	game := CreateGame().(*Game)

	game.Init([]byte(""), []string{"player1"})
	_, err := game.Play([]byte("{x: 0, y: 0}"), 1)
	fmt.Println(err)
	fmt.Println(game.GetState())
}
