package types

type Room struct {
	Name  string           `json:"name"`
	Users map[string]*User `json:"users"`
}
