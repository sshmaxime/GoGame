package server

import (
	"encoding/json"
	"github.com/GoGame/network"
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

func handleRequestInitParsing(w http.ResponseWriter, r *http.Request) (*network.InitRequest, error) {
	var req network.InitRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func handleRequestUpdateParsing(w http.ResponseWriter, r *http.Request) (*network.UpdateRequest, error) {
	var req network.UpdateRequest
	err := ReadBody(r.Body, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}
