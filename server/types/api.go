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
type ResponseGetUserByID struct {
	User *User `json:"user"`
}

// Rooms
type RequestPostRoom struct {
	ID string
}
type ResponsePostRoom struct {
	Room *Room
}

type ResponseGetAllRooms struct {
	Rooms map[string]*Room
}

type ResponseGetRoomByID struct {
	Room *Room
}

// Games
type ResponseGetAllGames struct {
	Games map[string]*GameDefinition
}

type ResponseGetGameByID struct {
	Game *GameDefinition
}

// Users
