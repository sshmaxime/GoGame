package models

type IGame interface {
	Init()

	Update(int)

	State() int
}
