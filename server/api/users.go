package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users map[string]*types.User
	var err error

	if users, err = database.GetAllUsers(); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, users)
	return
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user *types.User
	var err error

	ID, _ := mux.Vars(r)["id"]
	if user, err = database.GetUserByID(ID); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, user)
	return
}
