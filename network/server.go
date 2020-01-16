package network

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
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"games"`
}

type Server struct {
	handler *mux.Router

	// Configs
	NetworkConfig *NetworkConfig
	ServerConfig  *ServerConfig
}

func (s *Server) Init(networkConfig *NetworkConfig, serverConfig *ServerConfig) {
	s.NetworkConfig = networkConfig
	s.ServerConfig = serverConfig

	s.handler = mux.NewRouter()

	defaultRoutes := []ServerHandler{
		{Path: "/healthcheck", Fct: Healthcheck, Method: "GET"},
	}

	var gamesRoutes []ServerHandler
	for _, game := range serverConfig.Games {
		gameRoutes := []ServerHandler{
			{
				Path:   "/game/" + game.Name + "/init",
				Fct:    GameInit,
				Method: "GET",
			},
			{
				Path:   "/game/" + game.Name + "/update",
				Fct:    GameSendUpdate,
				Method: "POST",
			},
			{
				Path:   "/game/" + game.Name + "/update",
				Fct:    GameGetUpdate,
				Method: "GET",
			},
		}
		gamesRoutes = append(gamesRoutes, gameRoutes...)
	}

	routes := append(defaultRoutes, gamesRoutes...)
	for _, data := range routes {
		s.handler.HandleFunc(data.Path, data.Fct).Methods(data.Method)
	}
}

func (s *Server) Start() {
	err := http.ListenAndServe(":"+s.NetworkConfig.Port, s.handler)
	if err != nil {
		log.Println("Exiting the server ...")
	}
}

func (s *Server) Stop() {}
