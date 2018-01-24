package server

import (
	"net/http"

	"github.com/Crypto/cryptobookapp-rest-server/misc"
)

func (s *Server) DbStatus(w http.ResponseWriter, r *http.Request) {
	//url := r.URL.Query()
	dbStatus := GetDbStatus(s.Database)
	statusCode := http.StatusOK
	if !dbStatus.Healthy {
		statusCode = http.StatusInternalServerError
	}
	misc.WriteJson(w, &dbStatus, statusCode)
}

func (s *Server) Options(w http.ResponseWriter, r *http.Request) {
}
