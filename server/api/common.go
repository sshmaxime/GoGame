package api

import (
	"github.com/GoGame/database"
	"github.com/GoGame/types"
	"net/http"
)

// @Summary Auth admin
// @Description get admin info
// @Tags accounts,admin
// @Accept  json
// @Produce  json
// @Success 200 {object} types.RequestLogin
// @Failure 400 {object} types.Response
// @Security ApiKeyAuth
// @Router /admin/auth [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var request types.RequestLogin

	if err := readBody(r.Body, &request); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}

	user, err := database.AuthenticateUser(request.ID, request.Password)
	if err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var request types.RequestRegister

	if err := readBody(r.Body, &request); err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := database.CreateUser(request.ID, request.Password)
	if err != nil {
		responseAPI(w, http.StatusBadRequest, err)
		return
	}
	responseAPI(w, http.StatusOK, newUser)
}
