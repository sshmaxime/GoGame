package server

import (
	"github.com/GoGame/api"
	_ "github.com/GoGame/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title GoGame API
// @version 1.0
// @license.name Apache 2.0
// @description This is the documentation of GoGame. An open-source gaming server for small games.

func initAPI() error {
	server.handler = mux.NewRouter()

	// Login
	server.handler.HandleFunc("/api/login", api.Login).Methods(http.MethodPost)
	server.handler.HandleFunc("/api/register", api.Register).Methods(http.MethodPost)
	//

	// User
	server.handler.HandleFunc("/api/users", api.GetAllUsers).Methods(http.MethodGet)
	server.handler.HandleFunc("/api/users/{id}", api.GetUserByID).Methods(http.MethodGet)
	//

	// Room
	server.handler.HandleFunc("/api/rooms", api.PostRoom).Methods(http.MethodPost)

	server.handler.HandleFunc("/api/rooms", api.GetAllRooms).Methods(http.MethodGet)
	server.handler.HandleFunc("/api/rooms/{id}", api.GetRoomByID).Methods(http.MethodGet)
	//

	// Game
	server.handler.HandleFunc("/api/games", api.GetAllGames).Methods(http.MethodGet)
	server.handler.HandleFunc("/api/game/{id}", api.GetGameByID).Methods(http.MethodGet)
	//

	// Docs
	server.handler.PathPrefix("/api/doc").Handler(httpSwagger.WrapHandler)

	return nil
}
