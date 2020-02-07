package models

type IGame interface {
	Init()

	Play([]byte, string) error

	GetState() interface{}
}
