package models

// Game Packet
type InitRequest struct {
	GameName string `json:"game_name"`
}
type JoinRequest struct {
	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}
type StateRequest struct {
	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`
}
type UpdateRequest struct {
	GameName   string `json:"game_name"`
	GameRoomID string `json:"game_room_id"`

	Data interface{} `json:"data"`
}
