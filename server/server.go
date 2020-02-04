package server

import (
	"github.com/GoGame/managers"
	"github.com/GoGame/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	// Network handler
	handler *mux.Router

	// Config
	Config *models.ServerConfig

	// Management handlers
	ServerManager *managers.ServerManager
	AuthManager   *managers.AuthManager
	GameManager   *managers.GameManager
}

func (s *Server) Init(serverConfig *models.ServerConfig) error {
	s.Config = serverConfig

	s.initManagers()
	if err := s.initGames(); err != nil {
		return err
	}
	s.initHttp()
	return s.initRoutes()
}

func (s *Server) initGames() error {
	for _, game := range s.Config.Games {
		// Load the game engine creator function and store it
		gameEngineCreatorFunction, err := LoadGameEngineCreatorFunction(game.LibPath)
		if err != nil {
			return err
		}
		s.GameManager.StoreGame(game.Name, gameEngineCreatorFunction)
	}
	return nil
}

func (s *Server) initManagers() {
	s.ServerManager = new(managers.ServerManager)
	s.ServerManager.Init()

	s.AuthManager = new(managers.AuthManager)
	s.AuthManager.Init()

	s.GameManager = new(managers.GameManager)
	s.GameManager.Init()
}

func (s *Server) initHttp() {
	s.handler = mux.NewRouter()
}

func (s *Server) initRoutes() error {
	defaultRoutes := []Handler{
		{Path: "/healthcheck", Fct: s.Healthcheck, Method: "GET"},
	}

	gamesRoutes := []Handler{
		{
			Path: "/game/{gameName}/init", Fct: s.GameInit, Method: "POST",
		}, {
			Path: "/game/{gameName}/update", Fct: s.GameUpdate, Method: "POST",
		}, {
			Path: "/game/{gameName}/state", Fct: s.GameState, Method: "GET",
		},
	}

	// For each routes created, create its corresponding handleFunc
	for _, data := range append(defaultRoutes, gamesRoutes...) {
		s.handler.HandleFunc(data.Path, data.Fct).Methods(data.Method)
	}
	return nil
}

func (s *Server) Start() {
	log.Println("Server started on port: " + s.Config.Port)
	err := http.ListenAndServe(":"+s.Config.Port, s.handler)
	if err != nil {
		log.Println("Exiting the server on error ...", err)
	}
}
