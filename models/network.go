package models

// Game
type InitRequest struct {
	UserID   string `json:"user_id"`
	GameName string `json:"game_name"`
}
type JoinRequest struct {
	UserID string `json:"user_id"`

	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}
type UpdateRequest struct {
	UserID string `json:"user_id"`

	X          int    `json:"x"`
	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}

// Auth
type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
