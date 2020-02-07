package server

import (
	"encoding/json"
	"github.com/GoGame/models"
	"net/http"
)

func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	var request models.InitRequest
	user, err := s.GameGuard(r, &request)
	if err != nil {
		sendError(w, r, err)
		return
	}

	game, err := s.GameManager.CreateNewGame(request.GameName)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom := s.ServerManager.CreateGameRoom(request.GameName, game)
	if err := gameRoom.Join(user); err != nil {
		sendError(w, r, err)
		return
	}

	sendSuccessJSON(w, gameRoom)
}

func (s *Server) GameState(w http.ResponseWriter, r *http.Request) {
	var request models.StateRequest
	user, err := s.GameGuard(r, &request)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom, err := s.ServerManager.GetGameRoom(request.GameName, request.GameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	sendSuccessJSON(w, gameRoom.GetState(user.UserID))
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	var request models.UpdateRequest
	user, err := s.GameGuard(r, &request)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom, err := s.ServerManager.GetGameRoom(request.GameName, request.GameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	dataAsBytes, err := json.Marshal(request.Data)
	if err != nil {
		sendError(w, r, err)
		return
	}

	if err := gameRoom.Play(dataAsBytes, user.UserID); err != nil {
		sendError(w, r, err)
		return
	}
	sendSuccessJSON(w, nil)
}
