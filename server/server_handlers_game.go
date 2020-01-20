package server

import (
	"github.com/GoGame/network"
	"net/http"
)

func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	_, err := network.HandleRequestInitParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}
}

func (s *Server) GameJoin(w http.ResponseWriter, r *http.Request) {
	_, err := network.HandleRequestJoinParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	_, err := network.HandleRequestUpdateParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}
}

func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	gameRoomID := r.URL.Query().Get("gameRoomID")

	_, err := s.ServerManager.GetGameRoom(gameName, gameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}
}
