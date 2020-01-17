package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type NetworkConfig struct {
	Port string `yaml:"port"`
}

type ServerConfig struct {
	MaxSimultaneousPlayer uint `yaml:"max_simultaneous_player"`

	//
	Games []struct {
		Name    string `yaml:"name"`
		LibPath string `yaml:"lib_path"`
	} `yaml:"games"`
}

type Server struct {
	handler *mux.Router

	// Configs
	NetworkConfig *NetworkConfig
	ServerConfig  *ServerConfig

	// Game
	GameCreatorFunctions map[string]func() interface{}
}

func (s *Server) Init(networkConfig *NetworkConfig, serverConfig *ServerConfig) error {
	s.NetworkConfig = networkConfig
	s.ServerConfig = serverConfig

	s.handler = mux.NewRouter()
	s.GameCreatorFunctions = make(map[string]func() interface{})

	defaultRoutes := []Handler{
		{Path: "/healthcheck", Fct: s.Healthcheck, Method: "GET"},
	}

	var gamesRoutes []Handler
	for _, game := range serverConfig.Games {

		// Create routes
		gameRoutes := []Handler{
			{
				Path:   "/game/" + game.Name + "/init",
				Fct:    s.GameInit,
				Method: "GET",
			},
			{
				Path:   "/game/" + game.Name + "/update",
				Fct:    s.GameSendUpdate,
				Method: "POST",
			},
			{
				Path:   "/game/" + game.Name + "/update",
				Fct:    s.GameGetUpdate,
				Method: "GET",
			},
		}
		gamesRoutes = append(gamesRoutes, gameRoutes...)

		// Create game creator functions
		gameEngineCreatorFunction, err := LoadGameEngineCreatorFunction(game.LibPath)
		if err != nil {
			return err
		}

		s.GameCreatorFunctions[game.Name] = gameEngineCreatorFunction
	}

	routes := append(defaultRoutes, gamesRoutes...)
	for _, data := range routes {
		s.handler.HandleFunc(data.Path, data.Fct).Methods(data.Method)
	}

	return nil
}

func (s *Server) Start() {
	err := http.ListenAndServe(":"+s.NetworkConfig.Port, s.handler)
	if err != nil {
		log.Println("Exiting the server ...")
	}
}

func (s *Server) Stop() {}
