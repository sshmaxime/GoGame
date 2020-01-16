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
	Games struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}
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

	// TODO from config
	gamesRoutes := []ServerHandler{}
	//

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
