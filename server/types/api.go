package types

// Default API response
type ErrorResponse struct {
	Error string `json:"error"`
}

// Auth
type RequestLogin struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
type ResponseLogin struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

//
type RequestRegister struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
type ResponseRegister struct {
	User *User `json:"user"`
}

// Users
type ResponseGetAllUsers struct {
	Users map[string]*User `json:"users"`
}

//
type ResponseGetUserByID struct {
	User *User `json:"user"`
}

// Rooms
type ResponseGetAllRooms struct {
	Rooms map[string]*Room `json:"rooms"`
}

//
type ResponseGetRoomByID struct {
	Room *Room `json:"room"`
}

// Games
type ResponseGetAllGames struct {
	Games map[string]*GameDefinition `json:"games"`
}

//
type ResponseGetGameByID struct {
	Game *GameDefinition `json:"game"`
}
