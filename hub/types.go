package hub

import "github.com/GoGame/types"

const (
	MESSAGE_ROOM = "MESSAGE_ROOM"
	UPDATE_STATE = "UPDATE_STATE"
)

// Message
type Message struct {
	From string `json:"from"`
	Msg  string `json:"msg"`
}

// Info
type ServerState struct {
	OnlineUsers []*types.User `json:"online_users"`
	OnlineRooms []*types.Room `json:"online_rooms"`
}
