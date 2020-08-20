package database

import (
	"fmt"
	"github.com/GoGame/types"
)

type UserDatabase struct {
	users map[string]*types.User
}

var userDatabase UserDatabase

func InitUserDatabase() error {
	userDatabase.users = map[string]*types.User{
		"player1": {
			ID:       "player1",
			Password: "password",
		},
		"player2": {
			ID:       "player2",
			Password: "password",
		},
	}
	return nil
}

func AuthenticateUser(ID string, password string) (user *types.User, err error) {
	errorMsg := "error while authenticating user"

	if user, err = GetUserByID(ID); err != nil {
		return nil, fmt.Errorf("%v: %v", errorMsg, err)
	}
	if user.Password != password {
		return nil, fmt.Errorf("%v: %v", errorMsg, err)
	}
	return user, nil
}

func CreateUser(ID string, password string) (user *types.User, err error) {
	errorMsg := "error while creating user"

	if user = userDatabase.users[ID]; user != nil {
		return nil, fmt.Errorf("%v: user [%v] already exist", errorMsg, ID)
	}
	userDatabase.users[ID] = &types.User{
		ID:       ID,
		Password: password,
	}
	return userDatabase.users[ID], nil
}

func GetAllUsers() (user map[string]*types.User, err error) {
	return userDatabase.users, nil
}

func GetUserByID(ID string) (user *types.User, err error) {
	errorMsg := "error while getting user"

	if user = userDatabase.users[ID]; user == nil {
		return nil, fmt.Errorf("%v: user [%v] doesn't exist", errorMsg, ID)
	}
	return user, nil
}
