package models

type GameState struct {
	Board [][]rune
}

type Request struct {
}

type Response struct {
	GameState GameState
}

type IGame interface {
	Init(Request) Response

	Update(Request) Response

	State(Request) Response
}
