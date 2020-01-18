package main

type Game struct {
	x int
}

func (g *Game) Init() {
	g.x = 0
}

func (g *Game) Update(x int) {
	g.x = x
}

func (g *Game) State() int {
	return g.x
}

func CreateGame() interface{} {
	return new(Game)
}
