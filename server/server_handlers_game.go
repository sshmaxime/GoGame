package server

import (
	"encoding/json"
	"fmt"
	"github.com/GoGame/network"
	"net/http"
)

func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestInitParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	user, err := s.AuthManager.AuthenticateWithToken(req.UserID, req.UserToken)
	if err != nil {
		sendError(w, r, err)
		return
	}

	newGame, err := s.GameManager.CreateNewGame(req.GameName)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom := s.ServerManager.CreateGameRoom(req.GameName, newGame, user)
	sendSuccessJSON(w, gameRoom)
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	req, err := network.HandleRequestUpdateParsing(w, r)
	if err != nil {
		sendError(w, r, err)
		return
	}

	_, err = s.AuthManager.AuthenticateWithToken(req.UserID, req.UserToken)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom, err := s.ServerManager.GetGameRoom(req.GameName, req.GameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	fmt.Println("inside1")
	fmt.Println(req.Data)
	fmt.Println("inside2")
	bytess, _ := json.Marshal(req.Data)
	gameRoom.Game.Update(bytess, 1)
}

func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	gameRoomID := r.URL.Query().Get("gameRoomID")

	gameRoom, err := s.ServerManager.GetGameRoom(gameName, gameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	state := gameRoom.Game.State()
	sendSuccessJSON(w, state)
}
