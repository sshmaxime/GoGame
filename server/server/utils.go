package server

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"plugin"
)

func sendError(w http.ResponseWriter, _ *http.Request, err error) {
	http.Error(w, err.Error(), 500)
}
func sendSuccessJSON(w http.ResponseWriter, response interface{}) {
	responseAsJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	_, _ = w.Write(responseAsJSON)
}

func LoadGameEngineCreatorFunction(gameEngineLibPath string) (func() interface{}, error) {
	plug, err := plugin.Open(gameEngineLibPath)
	if err != nil {
		return nil, errors.New("impossible to find plugin")
	}

	symbol, err := plug.Lookup("CreateGame")
	if err != nil {
		return nil, errors.New("impossible to find symbol")
	}

	creatorFunc, ok := symbol.(func() interface{})
	if !ok {
		return nil, errors.New("unexpected type from module symbol")
	}

	return creatorFunc, nil
}

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
