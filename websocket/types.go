package websocket

const (
	LOGIN_REQUEST  = "LOGIN_REQUEST"
	LOGIN_RESPONSE = "LOGIN_RESPONSE"

	CREATE_ROOM_REQUEST  = "CREATE_ROOM_REQUEST"
	CREATE_ROOM_RESPONSE = "CREATE_ROOM_RESPONSE"

	JOIN_ROOM_REQUEST  = "JOIN_ROOM_REQUEST"
	JOIN_ROOM_RESPONSE = "JOIN_ROOM_RESPONSE"

	MESSAGE_ROOM_REQUEST  = "MESSAGE_ROOM_REQUEST"
	MESSAGE_ROOM_RESPONSE = "MESSAGE_ROOM_RESPONSE"

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
type MessageRoomRequest struct {
	RoomName string `json:"room_name"`
	Msg      string `json:"msg"`
}

// Response
type Response struct {
	Error string      `json:"error, omitempty"`
	Data  interface{} `json:"data, omitempty"`
}
