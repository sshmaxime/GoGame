package api

import (
	"encoding/json"
	"fmt"
	"github.com/GoGame/types"
	"io"
	"io/ioutil"
	"net/http"
)

func responseAPI(w http.ResponseWriter, status int, data interface{}) int {
	var response types.Response
	var responseAsBytes []byte
	var err error

	dataError, isError := data.(error)
	if isError {
		response.Error = dataError.Error()
	} else {
		response.Data = data
	}

	if responseAsBytes, err = json.Marshal(response); err != nil {
		fmt.Println("error occurred while parsing response object:" + err.Error())
		return writeResponseAPI(w, 500, nil)
	}
	return writeResponseAPI(w, status, responseAsBytes)
}

func writeResponseAPI(w http.ResponseWriter, status int, data []byte) int {
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		fmt.Println("Houston we got a problem: " + err.Error())
		return -1
	}
	return 0
}

func readBody(rBody io.Reader, output interface{}) error {
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
