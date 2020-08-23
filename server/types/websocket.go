package types

const (
	LOGIN_REQUEST  = "LOGIN_REQUEST"
	LOGIN_RESPONSE = "LOGIN_RESPONSE"

	CREATE_ROOM_REQUEST  = "CREATE_ROOM_REQUEST"
	CREATE_ROOM_RESPONSE = "CREATE_ROOM_RESPONSE"
)

// Login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	User User
}

// Create Room
type CreateRoomRequest struct {
	Name string `json:"name"`
}
type CreateRoomResponse struct {
}
