package server

import (
	"github.com/GoGame/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Server struct {
	handler *mux.Router
}

var server *Server

func Init() (err error) {
	server = &Server{}

	if err = initAPI(); err != nil {
		return err
	}

	if err = initDatabases(); err != nil {
		return err
	}

	if err = initWebSocket(); err != nil {
		return err
	}

	return nil
}

func Start() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	log.Println("Server started on port: " + config.GetPort())
	err := http.ListenAndServe("localhost:"+config.GetPort(), c.Handler(server.handler))
	if err != nil {
		log.Println("Exiting the server on error ...", err)
	}
}
