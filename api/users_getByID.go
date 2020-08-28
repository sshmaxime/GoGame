package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"github.com/gorilla/mux"
	"net/http"
)

// @Tags Users
// @Router /api/users/{id} [get]
///
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} ResponseGetUserByID
// @Failure 400 {object} ErrorResponse
func getUserByID(w http.ResponseWriter, r *http.Request) {
	var user *types.User
	var err error

	ID, _ := mux.Vars(r)["id"]
	if user, err = database.GetUserByID(ID); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, ResponseGetUserByID{
		User: user,
	})
	return
}
