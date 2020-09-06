package api

import (
	"github.com/GoGame/database"
	"net/http"
)

// @Tags Auth
// @Router /api/register [post]
///
// @Accept json
// @Param payload body RegisterRequest true "."
// @Produce json
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
func register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest

	if err := readBody(r.Body, &request); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := database.CreateUser(request.Username, request.Password)
	if err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, RegisterResponse{
		User: newUser,
	})
}
