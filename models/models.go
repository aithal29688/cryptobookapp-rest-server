package models

type ServerInfo struct {
	Server      string `json:"server"`
	Version     string `json:"version"`
	Hostname    string `json:"hostname"`
	Environment string `json:"environment"`
}

type DatabaseStatus struct {
	Healthy           bool   `json:"healthy"`
	Took              string `json:"took,omitempty"`
	Error             string `json:"error,omitempty"`
	OpenedConnections int    `json:"opened_connections"`
}

type DataRowD struct {
	Market 		string  `db:"market"`
	FromSymbol 	string 	`db:"fromsymbol"`
	ToSymbol	string 	`db:"tosymbol"`
	Price 		float64	`db:"price"`
	HighDay		float64	`db:"highday"`
	LowDay		float64	`db:"lowday"`
	OpenDay		float64	`db:"openday"`
	LastMarket	float64	`db:"lastmarket"`
	MarketCap	float64	`db:"marketcap"`
	Time		int64	`db:"time"`
}

type DataRowH struct {
	Market 		string
	FromSymbol 	string
	ToSymbol	string
	Price 		float64
	HighDay		float64
	LowDay		float64
	OpenDay		float64
	LastMarket	string
	MarketCap	float64
	Time		int64
}

type Status struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}
