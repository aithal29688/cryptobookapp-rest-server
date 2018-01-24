package misc

import (
	"encoding/json"
	"net/http"
	log "github.com/sirupsen/logrus"

	"github.com/Crypto/cryptobookapp-rest-server/models"
	"fmt"
)

func SendResponse(w http.ResponseWriter, success bool, reason string) {
	var status models.Status
	status.Success = success
	status.Reason = ""

	statusCode := http.StatusOK

	if !success {
		status.Reason = reason
		statusCode = http.StatusBadRequest

		log.WithField("reason", reason).Warn("Failed to process request")
	}

	WriteJson(w, status, statusCode)
}

func WriteText(w http.ResponseWriter, data string, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")

	w.WriteHeader(status)
	w.Write([]byte(data))
}

func WriteJson(w http.ResponseWriter, info interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Cache-Control", "no-store")

	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(info); err != nil {
		log.WithField("error", err).Info(fmt.Sprintf("Failed to write JSON"))
	}
}
