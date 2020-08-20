package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllGames(w http.ResponseWriter, _ *http.Request) {
	var users map[string]*types.GameDefinition
	var err error

	if users, err = database.GetAllGames(); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, users)
	return
}

func GetGameByID(w http.ResponseWriter, r *http.Request) {
	var user *types.GameDefinition
	var err error

	ID, _ := mux.Vars(r)["id"]
	if user, err = database.GetGameByID(ID); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, user)
	return
}
