package hub

const (
	MESSAGE_ROOM = "MESSAGE_ROOM"
)

// Message
type Message struct {
	From string `json:"from"`
	Msg  string `json:"msg"`
}
