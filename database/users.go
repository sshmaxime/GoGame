package database

import (
	"errors"
	"fmt"
	"github.com/GoGame/types"
)

type UserDatabase struct {
	users map[string]*types.User // map[username]password
	auth  map[string]string      // map[username]password
}

var userDatabase UserDatabase

func initUserDatabase() error {
	userDatabase.users = map[string]*types.User{
		"player1": {
			Username: "player1",
		},
		"player2": {
			Username: "player2",
		},
	}
	userDatabase.auth = map[string]string{
		"player1": "player1",
		"player2": "player2",
	}
	return nil
}

func AuthenticateUser(Username string, password string) (user *types.User, err error) {
	errorMsg := "error while authenticating user " + Username

	var passwd string
	var ok bool

	if passwd, ok = userDatabase.auth[Username]; ok == false {
		return nil, fmt.Errorf("%v: %v", errorMsg, errors.New("user doesn't exist"))
	}

	if passwd != password {
		return nil, fmt.Errorf("%v: %v", errorMsg, errors.New("bad credential"))
	}
	return userDatabase.users[Username], nil
}

func CreateUser(username string, password string) (user *types.User, err error) {
	errorMsg := "error while creating user"

	if user = userDatabase.users[username]; user != nil {
		return nil, fmt.Errorf("%v: user [%v] already exist", errorMsg, username)
	}
	userDatabase.users[username] = &types.User{
		Username: username,
	}
	userDatabase.auth[username] = password
	return userDatabase.users[username], nil
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
