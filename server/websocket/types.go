package websocket

const (
	LOGIN_REQUEST = "LOGIN_REQUEST"
	LOGIN_SUCCESS = "LOGIN_SUCCESS"

	REGISTER_REQUEST = "REGISTER_REQUEST"
	REGISTER_SUCCESS = "REGISTER_SUCCESS"

	CREATE_ROOM_REQUEST = "CREATE_ROOM_REQUEST"
	CREATE_ROOM_SUCCESS = "CREATE_ROOM_SUCCESS"

	JOIN_ROOM_REQUEST = "JOIN_ROOM_REQUEST"
	JOIN_ROOM_SUCCESS = "JOIN_ROOM_SUCCESS"

	LEAVE_ROOM_REQUEST = "LEAVE_ROOM_REQUEST"
	LEAVE_ROOM_SUCCESS = "LEAVE_ROOM_SUCCESS"

	MESSAGE_ROOM_REQUEST = "MESSAGE_ROOM_REQUEST"
	MESSAGE_ROOM_SUCCESS = "MESSAGE_ROOM_SUCCESS"

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

// Response
type Response struct {
	Error string      `json:"error, omitempty"`
	Data  interface{} `json:"data, omitempty"`
}
