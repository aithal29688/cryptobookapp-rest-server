package misc

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v1"
)

type Http struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
}

type Database struct {
	Host            string
	Username        string
	Password        string
	DbName          string
	Port            int
	Ssl             bool
	DbCreate        bool
	ConnectTimeout  int
	ValidTables     []string
	NumMaxIdleConns int
	NumMaxOpenConns int
}

type Log struct {
	Level       string
	Filename    string
	MaxSizeMB   int
	MaxBackups  int
	MaxAgeDays  int
	WriteStdout bool
	Json        bool
}

type Conf struct {
	Http       Http
	Database   Database
	StatsdHost string
	Env        string
	Log        Log
}

func LoadConf(filename string, appName string) (*Conf, error) {
	cnf := Conf{}

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(source, &cnf); err != nil {
		return nil, err
	}

	log.Info("Loaded config file")

	return &cnf, nil
}
