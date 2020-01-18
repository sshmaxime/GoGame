package network

import (
	"encoding/json"
	"github.com/GoGame/models"
	"io"
	"io/ioutil"
	"net/http"
)

func ReadBody(rBody io.Reader, output interface{}) error {
	body, err := ioutil.ReadAll(rBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}

func HandleRequestInitParsing(w http.ResponseWriter, r *http.Request) (*models.InitRequest, error) {
	var req models.InitRequest
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
