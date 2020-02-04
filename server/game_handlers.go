package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoGame/models"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) validateGameActionRequest(r *http.Request, request interface{}) (*models.User, string, error) {
	gameName, ok := mux.Vars(r)["gameName"]
	if !ok {
		return nil, "", errors.New("cannot find gameName in request")
	}

	token := r.Header.Get("X-Session-Token")
	user := s.AuthManager.AuthenticateWithToken(token)
	if user == nil {
		return nil, "", fmt.Errorf("cannot find user with token [%v]", token)
	}

	err := ReadBody(r.Body, &request)
	if err != nil {
		return nil, "", fmt.Errorf("cannot parse body request")
	}

	return user, gameName, nil
}

func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	var request models.InitRequest
	user, gameName, err := s.validateGameActionRequest(r, &request)
	if err != nil {
		sendError(w, r, err)
		return
	}

	newGame, err := s.GameManager.CreateNewGame(gameName)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom := s.ServerManager.CreateGameRoom(gameName, newGame, user)
	sendSuccessJSON(w, gameRoom)
}

func (s *Server) GameUpdate(w http.ResponseWriter, r *http.Request) {
	var request models.UpdateRequest
	_, gameName, err := s.validateGameActionRequest(r, &request)
	if err != nil {
		sendError(w, r, err)
		return
	}

	gameRoom, err := s.ServerManager.GetGameRoom(gameName, request.GameRoomID)
	if err != nil {
		sendError(w, r, err)
		return
	}

	bytess, _ := json.Marshal(request.Data)
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
