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

type Status struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
}
