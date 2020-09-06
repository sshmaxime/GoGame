package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/GoGame/docs"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"io"
	"io/ioutil"
	"net/http"
)

var CorsMiddleware = cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"GET", "POST"},
})

// @title GoGame API
// @version 1.0
// @license.name Apache 2.0
// @description This is the documentation of GoGame. An open-source gaming server for small games.
func Init() (handler *mux.Router, err error) {
	handler = mux.NewRouter()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello world :)"))
	})

	// Register
	handler.HandleFunc("/api/register", register).Methods(http.MethodPost)
	//

	// User
	handler.HandleFunc("/api/users", getAllUsers).Methods(http.MethodGet)
	handler.HandleFunc("/api/users/{id}", getUserByID).Methods(http.MethodGet)
	//

	// Game
	handler.HandleFunc("/api/games", getAllGames).Methods(http.MethodGet)
	handler.HandleFunc("/api/game/{id}", getGameByID).Methods(http.MethodGet)
	//

	// Documentation
	handler.PathPrefix("/api/doc").Handler(httpSwagger.WrapHandler)

	return handler, nil
}

func errorAPI(w http.ResponseWriter, status int, err error) int {
	var response ErrorResponse
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
