package server

import (
	"net/http"

	"github.com/Crypto/cryptobookapp-rest-server/misc"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/Crypto/cryptobookapp-rest-server/models"
)

var(
	market string
	fromsymbol string
	tosymbol string
	price float64
	lastupdate int64
	openday float64
	highday float64
	lowday float64
	lastmarket string
	marketcap float64
)

const (
	from_symbol = "fromsymbol"
	from_time = "fromTime"
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

func (s *Server) ExtractData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fromSymbol := vars[from_symbol]
	fromTime := vars[from_time]

	rows, err := s.Database.Query("select market,fromsymbol,tosymbol,price,lastupdate,openday,highday,lowday,lastmarket,mktcap from cryptoprice where fromsymbol = $1 and lastupdate > $2 and tosymbol = 'USD'", fromSymbol, fromTime)
	if err != nil {
		log.Fatal(err)
	}
	dataRows := []models.DataRowH{}
	for rows.Next() {
		err := rows.Scan(&market,&fromsymbol,&tosymbol,&price,&lastupdate,&openday,&highday,&lowday,&lastmarket,&marketcap)
		if err != nil {
			log.Fatal(err)
		}
		dataRow := models.DataRowH{Market:market,FromSymbol:fromsymbol,ToSymbol:tosymbol,Price:price,Time:lastupdate,OpenDay:openday,HighDay:highday,LowDay:lowday,LastMarket:lastmarket,MarketCap:marketcap}
		dataRows = append(dataRows, dataRow)
	}

	statusCode := http.StatusOK
	misc.WriteJson(w, &dataRows, statusCode)
	defer rows.Close()
}
