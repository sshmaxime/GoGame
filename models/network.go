package models

type UserIdentification struct {
	UserID    string `json:"user_id"`
	UserToken string `json:"user_token"`
}

// Game
type InitRequest struct {
	UserIdentification

	GameName string `json:"game_name"`
}
type JoinRequest struct {
	UserIdentification

	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}
type UpdateRequest struct {
	UserIdentification

	Data       interface{} `json:"data"`
	GameName   string      `json:"game_name"`
	GameRoomID string      `json:"game_room_id"`
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
