package api

import "github.com/GoGame/types"

// Default API response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Register
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterResponse struct {
	User *types.User `json:"user"`
}

// Users
type GetAllUsersResponse struct {
	Users map[string]*types.User `json:"users"`
}

//
type ResponseGetUserByID struct {
	User *types.User `json:"user"`
}

// Games
type GetAllGamesResponse struct {
	Games map[string]*types.GameDefinition `json:"games"`
}

//
type GetGameByIDResponse struct {
	Game *types.GameDefinition `json:"game"`
}
