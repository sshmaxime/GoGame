package main

type Game struct {
	x int
}

func (g *Game) Init() {
	g.x = 0
}

func (g *Game) Add() {
	g.x = g.x + 1
}

func (g *Game) Get() int {
	return g.x
}

func CreateGame() interface{} {
	return new(Game)
}
