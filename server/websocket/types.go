package websocket

const (
	LOGIN_REQUEST    = "LOGIN_REQUEST"
	LOGIN_SUCCESS    = "LOGIN_SUCCESS"
	REGISTER_REQUEST = "REGISTER_REQUEST"
	REGISTER_SUCCESS = "REGISTER_SUCCESS"

	// Room
	CREATE_ROOM_REQUEST  = "CREATE_ROOM_REQUEST"
	CREATE_ROOM_SUCCESS  = "CREATE_ROOM_SUCCESS"
	JOIN_ROOM_REQUEST    = "JOIN_ROOM_REQUEST"
	JOIN_ROOM_SUCCESS    = "JOIN_ROOM_SUCCESS"
	LEAVE_ROOM_REQUEST   = "LEAVE_ROOM_REQUEST"
	LEAVE_ROOM_SUCCESS   = "LEAVE_ROOM_SUCCESS"
	MESSAGE_ROOM_REQUEST = "MESSAGE_ROOM_REQUEST"
	MESSAGE_ROOM_SUCCESS = "MESSAGE_ROOM_SUCCESS"

	// Game
	CREATE_GAME_REQUEST = "CREATE_GAME_REQUEST"
	CREATE_GAME_SUCCESS = "CREATE_GAME_SUCCESS"
	JOIN_GAME_REQUEST   = "JOIN_GAME_REQUEST"
	JOIN_GAME_SUCCESS   = "JOIN_GAME_SUCCESS"
	LEAVE_GAME_REQUEST  = "LEAVE_GAME_REQUEST"
	LEAVE_GAME_SUCCESS  = "LEAVE_GAME_SUCCESS"
	PLAY_GAME_REQUEST   = "PLAY_GAME_REQUEST"
	PLAY_GAME_SUCCESS   = "PLAY_GAME_SUCCESS"

	ERROR = "ERROR"
)

// Login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Room requests
type CreateRoomRequest struct {
	Name string `json:"name"`
}
type JoinRoomRequest struct {
	Name string `json:"name"`
}
type LeaveRoomRequest struct {
	Name string `json:"name"`
}
type MessageRoomRequest struct {
	RoomName string `json:"room_name"`
	Msg      string `json:"msg"`
}

// Game requests
type CreateGameRequest struct {
	RoomName string `json:"room_name"`
	Name     string `json:"name"`
}
type JoinGameRequest struct {
	RoomName string `json:"room_name"`
	Name     string `json:"name"`
}
type LeaveGameRequest struct {
	RoomName string `json:"room_name"`
	Name     string `json:"name"`
}
type PlayGameRequest struct {
	RoomName string `json:"room_name"`
	Data     string `json:"data"`
}

// Response
type Response struct {
	Error string      `json:"error, omitempty"`
	Data  interface{} `json:"data, omitempty"`
}
