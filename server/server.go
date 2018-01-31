package server

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/Crypto/cryptobookapp-rest-server/misc"
	"github.com/Crypto/cryptobookapp-rest-server/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Info     *models.ServerInfo
	Uptime   time.Time
	Database *sql.DB
	Server   *http.Server
}

func (s *Server) Run(conf *misc.Http, address string, routing Route) error {

	s.Server = &http.Server{
		Addr:              address,
		ReadTimeout:       time.Duration(conf.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(conf.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(conf.WriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(conf.ReadTimeout*2) * time.Second,
		Handler:           mux.NewRouter().StrictSlash(true),
	}

	routing(s)

	log.WithField("address", address[1:]).Info("HTTP server ready")

	return s.Server.ListenAndServe()
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Server.Shutdown(ctx); err != nil {
		//logging.Logger.WithField("error", err).Warn("Failed to shut down http server cleanly")

		// Close all open connections
		if err := s.Server.Close(); err != nil {
			//logging.Logger.WithField("error", err).Warn("Failed to force-close http server")
		}
	}

	return nil
}
