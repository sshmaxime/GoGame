package types

// Default API response
type Response struct {
	Error string      `json:"error, omitempty"`
	Data  interface{} `json:"data, omitempty"`
}

// Common
type RequestLogin struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type RequestRegister struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// Rooms
type RequestPostRoom struct {
	ID string
}

// Games

// Users
