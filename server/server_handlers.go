package server

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Path   string
	Fct    func(w http.ResponseWriter, r *http.Request)
	Method string
}

// Default routes
func (s *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// Utils
func sendError(w http.ResponseWriter, r *http.Request, err error) {
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
