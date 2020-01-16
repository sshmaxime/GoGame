package network

import "net/http"

type ServerHandler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func GameInit(w http.ResponseWriter, r *http.Request)       {}
func GameSendUpdate(w http.ResponseWriter, r *http.Request) {}
func GameGetUpdate(w http.ResponseWriter, r *http.Request)  {}
