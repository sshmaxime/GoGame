package network

import (
	"github.com/GoGame/models"
	"net/http"
)

func HandleRequestInitParsing(w http.ResponseWriter, r *http.Request) (*models.InitRequest, error) {
	var req models.InitRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func HandleRequestJoinParsing(w http.ResponseWriter, r *http.Request) (*models.JoinRequest, error) {
	var req models.JoinRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func HandleRequestUpdateParsing(w http.ResponseWriter, r *http.Request) (*models.UpdateRequest, error) {
	var req models.UpdateRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
