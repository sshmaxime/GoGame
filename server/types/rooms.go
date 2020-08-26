package types

type Room struct {
	Name  string  `json:"name"`
	Users []*User `json:"users"`
}
