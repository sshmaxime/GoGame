package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

// @Tags Rooms
// @Router /api/rooms [post]
///
// @Produce json
// @Success 200 {object} types.ResponsePostRoom
// @Failure 400 {object} types.ErrorResponse
func PostRoom(w http.ResponseWriter, r *http.Request) {
	var request types.RequestPostRoom

	if err := readBody(r.Body, &request); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}

	room, err := database.CreateRoom(request.ID)
	if err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponsePostRoom{
		Room: room,
	})
}

// @Tags Rooms
// @Router /api/rooms [get]
///
// @Produce json
// @Success 200 {object} types.ResponseGetAllRooms
// @Failure 400 {object} types.ErrorResponse
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	var rooms map[string]*types.Room
	var err error

	if rooms, err = database.GetAllRooms(); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseGetAllRooms{
		Rooms: rooms,
	})
	return
}

// @Tags Rooms
// @Router /api/rooms/{id} [get]
///
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} types.ResponseGetRoomByID
// @Failure 400 {object} types.ErrorResponse
func GetRoomByID(w http.ResponseWriter, r *http.Request) {
	var room *types.Room
	var err error

	ID, _ := mux.Vars(r)["id"]
	if room, err = database.GetRoomByID(ID); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseGetRoomByID{
		Room: room,
	})
	return
}
