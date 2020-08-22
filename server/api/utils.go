package api

import (
	"encoding/json"
	"fmt"
	"github.com/GoGame/types"
	"io"
	"io/ioutil"
	"net/http"
)

func errorAPI(w http.ResponseWriter, status int, err error) int {
	var response types.ErrorResponse
	var responseAsBytes []byte
	var errTmp error

	response.Error = err.Error()
	if responseAsBytes, errTmp = json.Marshal(response); errTmp != nil {
		fmt.Println("error occurred while parsing error response object:" + errTmp.Error())
		return writeResponseAPI(w, 500, nil)
	}
	return writeResponseAPI(w, status, responseAsBytes)
}

func successAPI(w http.ResponseWriter, status int, data interface{}) int {
	var responseAsBytes []byte
	var err error

	if responseAsBytes, err = json.Marshal(data); err != nil {
		fmt.Println("error occurred while parsing success response object:" + err.Error())
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
