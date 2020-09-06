package server

import (
	"github.com/GoGame/api"
	"github.com/GoGame/config"
	"github.com/GoGame/database"
	"github.com/GoGame/websocket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type server struct {
	handler *mux.Router

	wsHandler *socketio.Server
}

func Start() (err error) {
	server := &server{}

	if server.handler, err = api.Init(); err != nil {
		return err
	}

	if server.wsHandler, err = websocket.Init(); err != nil {
		return err
	}
	server.handler.Handle("/socket.io/", websocket.CorsMiddleware(server.wsHandler))

	if err = database.Init(); err != nil {
		return err
	}

	go server.wsHandler.Serve()
	defer server.wsHandler.Close()

	log.Println("Server started on: [" + config.GetAddress() + ":" + config.GetPort() + "]")
	if err = http.ListenAndServe(config.GetAddress()+":"+config.GetPort(), api.CorsMiddleware.Handler(server.handler)); err != nil {
		log.Println("Exiting the server, error: ", err)
	}
	return nil
}
