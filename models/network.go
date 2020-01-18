package models

type InitRequest struct {
	UserID string `json:"user_id"`
}

type UpdateRequest struct {
	UserID string `json:"user_id"`
}
