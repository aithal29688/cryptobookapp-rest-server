package server

import (
	"database/sql"
	"time"

	"github.com/Crypto/cryptobookapp-rest-server/models"
)

func GetDbStatus(db *sql.DB) models.DatabaseStatus {
	dbStatus := models.DatabaseStatus{Healthy: false}
	if db != nil {
		if err := db.Ping(); err == nil {
			t1 := time.Now()
			rows, err := db.Query("SELECT 1")
			if err != nil {
				dbStatus.Healthy = false
				dbStatus.Error = err.Error()
				dbStatus.Took = time.Since(t1).String()
			} else {
				dbStatus.Took = time.Since(t1).String()
				dbStatus.Healthy = true
				dbStatus.OpenedConnections = db.Stats().OpenConnections
			}

			defer rows.Close()
		} else {
			dbStatus.Error = err.Error()
		}
	}
	return dbStatus
}
