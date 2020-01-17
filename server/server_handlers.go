package server

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

var game IGame

func (s *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (s *Server) GameInit(w http.ResponseWriter, r *http.Request) {
	game, _ = s.GetNewGame("tictactoe")
	game.Init()
}
func (s *Server) GameSendUpdate(w http.ResponseWriter, r *http.Request) {
	game.Add()
}
func (s *Server) GameGetUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println(game.Get())
}
