package api

import (
	_ "github.com/GoGame/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title GoGame API
// @version 1.0
// @license.name Apache 2.0
// @description This is the documentation of GoGame. An open-source gaming server for small games.
func Init() (handler *mux.Router, err error) {
	handler = mux.NewRouter()

	// Register
	handler.HandleFunc("/api/register", register).Methods(http.MethodPost)
	//

	// User
	handler.HandleFunc("/api/users", getAllUsers).Methods(http.MethodGet)
	handler.HandleFunc("/api/users/{id}", getUserByID).Methods(http.MethodGet)
	//

	// Game
	handler.HandleFunc("/api/games", getAllGames).Methods(http.MethodGet)
	handler.HandleFunc("/api/game/{id}", getGameByID).Methods(http.MethodGet)
	//

	// Documentation
	handler.PathPrefix("/api/doc").Handler(httpSwagger.WrapHandler)

	return handler, nil
}
