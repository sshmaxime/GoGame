package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

func PostRoom(w http.ResponseWriter, r *http.Request) {
	var request types.RequestPostRoom

	if err := readBody(r.Body, &request); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}

	room, err := database.CreateRoom(request.ID)
	if err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, room)
}

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var users map[string]*types.Room
	var err error

	if users, err = database.GetAllRooms(); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, users)
	return
}

func GetRoomByID(w http.ResponseWriter, r *http.Request) {
	var user *types.Room
	var err error

	ID, _ := mux.Vars(r)["id"]
	if user, err = database.GetRoomByID(ID); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, user)
	return
}
