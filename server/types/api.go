package types

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
	User *User `json:"user"`
}

// Users
type GetAllUsersResponse struct {
	Users map[string]*User `json:"users"`
}

//
type ResponseGetUserByID struct {
	User *User `json:"user"`
}

// Games
type GetAllGamesResponse struct {
	Games map[string]*GameDefinition `json:"games"`
}

//
type GetGameByIDResponse struct {
	Game *GameDefinition `json:"game"`
}
