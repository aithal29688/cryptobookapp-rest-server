package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Crypto/cryptobookapp-rest-server/misc"
	"github.com/Crypto/cryptobookapp-rest-server/models"
	"github.com/Crypto/cryptobookapp-rest-server/server"
	log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

var (
	showVersion = flag.Bool("version", false, "print version string")
	configFile  = flag.String("config", "dev.config.yaml", "Config file to use")
	hostname    = "unknown"

	Version    = "unknown"
	GitRefs    = "unknown"
	GitVersion = "unknown"
)

func main() {
	flag.Parse()
	conf, err := misc.LoadConf(*configFile, "cryptobookapp-rest-service")
	if err != nil {
		log.WithField("error", err).Error("Failed to load config")
		os.Exit(-1)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	if host, err := os.Hostname(); err == nil {
		hostname = host
	}

	db, err := newDatabaseSession(&conf.Database)
	if err != nil {
		log.WithField("error", err).Error("Failed to open database")
		shutdown(nil, nil)
	}

	s := startHTTPServer(conf.Http.Address, server.AddRoutes, db, conf, termChan)

	<-termChan

	shutdown(s, db)
}

func startHTTPServer(address string, routing server.Route, db *sql.DB, conf *misc.Conf, termChan chan os.Signal) *server.Server {
	hostname, _ := os.Hostname()
	s := &server.Server{
		Info: &models.ServerInfo{
			Server:      "cryptobookapp-rest-server",
			Version:     Version,
			Hostname:    hostname,
			Environment: conf.Env,
		},
		Uptime:   time.Now(),
		Database: db,
	}

	go func() {
		if err := s.Run(&conf.Http, address, routing); err != nil && err != http.ErrServerClosed {
			log.WithField("error", err).Error("Failed to start HTTP server")

		}

		termChan <- syscall.SIGTERM
	}()

	return s
}

func newDatabaseSession(conf *misc.Database) (*sql.DB, error) {
	log.WithField("config", *conf).Info("Connecting to the database...")
	dbParams := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d sslmode=disable", conf.Host, conf.DbName,
		conf.Username, conf.Password, conf.Port)

	return sql.Open("postgres", dbParams)
}

func shutdown(s *server.Server, db *sql.DB) {

	if s != nil {
		s.Close()
	}

	if db != nil {
		db.Close()
	}

	log.Info("It was nice talking with you. Hope to see you later!")
	os.Exit(0)
}
