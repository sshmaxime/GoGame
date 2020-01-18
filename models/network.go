package models

type InitRequest struct {
	UserID   string `json:"user_id"`
	GameName string `json:"game_name"`
}

type UpdateRequest struct {
	UserID string `json:"user_id"`

	X          int    `json:"x"`
	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}
