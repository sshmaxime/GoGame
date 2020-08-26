package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"net/http"
)

// @Tags Users
// @Router /api/users [get]
///
// @Produce json
// @Success 200 {object} GetAllUsersResponse
// @Failure 400 {object} ErrorResponse
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users map[string]*types.User
	var err error

	if users, err = database.GetAllUsers(); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, GetAllUsersResponse{
		Users: users,
	})
	return
}
