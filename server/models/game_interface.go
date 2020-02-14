package models

type IGame interface {
	Init()

	Play([]byte, uint8) error

	GetState() interface{}
}
