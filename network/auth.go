package network

import (
	"github.com/GoGame/models"
	"net/http"
)

func HandleRequestLoginParsing(w http.ResponseWriter, r *http.Request) (*models.LoginRequest, error) {
	var req models.LoginRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func HandleRequestRegisterParsing(w http.ResponseWriter, r *http.Request) (*models.RegisterRequest, error) {
	var req models.RegisterRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
