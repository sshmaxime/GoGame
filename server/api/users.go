package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

// @Tags Users
// @Router /api/users [get]
///
// @Produce json
// @Success 200 {object} types.ResponseGetAllUsers
// @Failure 400 {object} types.ErrorResponse
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users map[string]*types.User
	var err error

	if users, err = database.GetAllUsers(); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseGetAllUsers{
		Users: users,
	})
	return
}

// @Tags Users
// @Router /api/users/{id} [get]
///
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} types.ResponseGetUserByID
// @Failure 400 {object} types.ErrorResponse
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var user *types.User
	var err error

	ID, _ := mux.Vars(r)["id"]
	if user, err = database.GetUserByID(ID); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseGetUserByID{
		User: user,
	})
	return
}
