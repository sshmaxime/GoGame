package server

import (
	"fmt"
	"github.com/GoGame/network"
	"net/http"
)

type Handler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

// Default routes
func (s *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// Game routes
func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestInitParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	user := s.AuthManager.GetUser(req.UserID)
	if user == nil {
		sendError(w, r, err)
		return
	}

	newGame, err := s.GameManager.GetGame(req.GameName)
	if err != nil {
		sendError(w, r, err)
		return
	}

	s.ServerManager.CreateGameRoom(req.GameName, newGame)
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestUpdateParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom, err := s.ServerManager.GetGameRoom(req.GameName, req.GameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom.Game.Update(req.X)
}
func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	gameRoomID := r.URL.Query().Get("gameRoomID")

	gameRoom, err := s.ServerManager.GetGameRoom(gameName, gameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	fmt.Println(gameRoom.Game.State())
}

func sendError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), 500)
}
