package server

import (
	"github.com/GoGame/managers"
	"github.com/GoGame/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	handler *mux.Router

	// Configs
	Config *models.ServerConfig

	// Management
	ServerManager *managers.ServerManager
	AuthManager   *managers.AuthManager
	GameManager   *managers.GameManager
}

func (s *Server) Init(serverConfig *models.ServerConfig) error {
	s.handler = mux.NewRouter()
	s.Config = serverConfig

	s.ServerManager = new(managers.ServerManager)
	s.ServerManager.Init()

	s.AuthManager = new(managers.AuthManager)
	s.AuthManager.Init()

	s.GameManager = new(managers.GameManager)
	s.GameManager.Init()

	return s.initRoutes()
}

func (s *Server) initRoutes() error {
	defaultRoutes := []Handler{
		{Path: "/healthcheck", Fct: s.Healthcheck, Method: "GET"},
	}

	var gamesRoutes []Handler
	for _, game := range s.Config.Games {

		// Create routes
		gameRoutes := []Handler{
			{
				Path:   "/game/" + game.Name + "/init",
				Fct:    s.GameInit,
				Method: "POST",
			},
			{
				Path:   "/game/" + game.Name + "/update",
				Fct:    s.GameUpdate,
				Method: "POST",
			},
			{
				Path:   "/game/" + game.Name + "/state",
				Fct:    s.GameState,
				Method: "GET",
			},
		}
		gamesRoutes = append(gamesRoutes, gameRoutes...)

		gameEngineCreatorFunction, err := LoadGameEngineCreatorFunction(game.LibPath)
		if err != nil {
			return err
		}

		s.GameManager.AddGame(game.Name, gameEngineCreatorFunction)
	}

	routes := append(defaultRoutes, gamesRoutes...)
	for _, data := range routes {
		s.handler.HandleFunc(data.Path, data.Fct).Methods(data.Method)
	}

	return nil
}

func (s *Server) Start() {
	log.Println("Server started on port: " + s.Config.Port)
	err := http.ListenAndServe(":"+s.Config.Port, s.handler)
	if err != nil {
		log.Println("Exiting the server ...")
	}
}
