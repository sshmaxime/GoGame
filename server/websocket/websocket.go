package websocket

import (
	"fmt"
	"github.com/GoGame/hub"
	"github.com/googollee/go-socket.io"
	"net/http"
)

func Init() (wsHandler *socketio.Server, err error) {

	if wsHandler, err = hub.Init(); err != nil {
		return nil, err
	}

	wsHandler.OnConnect("/", func(socket socketio.Conn) error {
		fmt.Println("onconnect:" + socket.ID())
		return nil
	})
	wsHandler.OnDisconnect("/", func(socket socketio.Conn, str string) {
		// TODO
	})

	initAuth(wsHandler)
	initRooms(wsHandler)
	initInfo(wsHandler)

	return wsHandler, nil
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("Origin")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}
