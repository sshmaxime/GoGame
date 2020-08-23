package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"net/http"
)

// @Tags Auth
// @Router /api/register [post]
///
// @Accept json
// @Param payload body types.RegisterRequest true "."
// @Produce json
// @Success 200 {object} types.RegisterResponse
// @Failure 400 {object} types.ErrorResponse
func register(w http.ResponseWriter, r *http.Request) {
	var request types.RegisterRequest

	if err := readBody(r.Body, &request); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := database.CreateUser(request.Username, request.Password)
	if err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.RegisterResponse{
		User: newUser,
	})
}
