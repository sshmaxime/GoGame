package types

// Interface similar to every game
type IGame interface {
	Init([]byte, []string)

	Play([]byte, string) (interface{}, error)

	GetState() interface{}
}

type GameDefinition struct {
	ID              string `yaml:"id"`
	MaxPlayer       int    `yaml:"max_player"`
	LibPath         string `yaml:"lib_path"`
	CreatorFunction func() interface{}
}

type Game struct {
	GameDefinition *GameDefinition

	Game  IGame            `yaml:",omitempty"`
	Users map[string]*User `json:"users"`
}
