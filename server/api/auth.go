package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"net/http"
)

// @Tags Auth
// @Router /api/login [post]
///
// @Accept json
// @Param payload body types.RequestLogin true "."
// @Produce json
// @Success 200 {object} types.ResponseLogin
// @Failure 400 {object} types.ErrorResponse
func login(w http.ResponseWriter, r *http.Request) {
	var request types.RequestLogin

	if err := readBody(r.Body, &request); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}

	user, err := database.AuthenticateUser(request.ID, request.Password)
	if err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseLogin{
		User:  user,
		Token: "nothing",
	})
}

// @Tags Auth
// @Router /api/register [post]
///
// @Accept json
// @Param payload body types.RequestRegister true "."
// @Produce json
// @Success 200 {object} types.ResponseRegister
// @Failure 400 {object} types.ErrorResponse
func register(w http.ResponseWriter, r *http.Request) {
	var request types.RequestRegister

	if err := readBody(r.Body, &request); err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := database.CreateUser(request.ID, request.Password)
	if err != nil {
		errorAPI(w, http.StatusBadRequest, err)
		return
	}
	successAPI(w, http.StatusOK, types.ResponseRegister{
		User: newUser,
	})
}
