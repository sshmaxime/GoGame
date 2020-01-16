package network

import "net/http"

type ServerHandler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
}
