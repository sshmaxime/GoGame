package models

type IGame interface {
	Init()

	Update([]byte, uint8)

	State() interface{}
}
