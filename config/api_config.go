package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"strconv"
)

type APIConf struct {
	Env    string
	Debug  bool
	Server struct {
		Hostname string
		Port     int
	}
	Database struct {
		Url string
	}
	Migration struct {
		Path string
	}
	Jwt struct {
		SignKey string
	}
	Networks []string
}

var APIConfig APIConf

func ParseAPIConf(cfg string) (*APIConf, error) {
	conf := APIConf{}

	if data, err := ioutil.ReadFile(cfg); err != nil {
		log.Fatalf("Could not read config file:%s because of %s.", cfg, err)
	} else {
		if _, err := toml.Decode(string(data), &conf); err != nil {
			return nil, errors.New("Could not read TOML config")
		}
	}

	dbUrl := getEnv("DATABASE_URL", "postgres:5432")
	dbUser := getEnv("DATABASE_USERNAME", "user")
	dbPass := getEnv("DATABASE_PASSWORD", "pass")
	dbTable := getEnv("DATABASE_TABLE", "tezoslink?sslmode=disable")
	dbParam := getEnv("DATABASE_ADDITIONAL_PARAMETER", "sslmode=disable")
	conf.Database.Url = fmt.Sprintf("postgres://%s:%s@%s/%s?%s", dbUser, dbPass, dbUrl, dbTable, dbParam)

	conf.Server.Hostname = getEnv("SERVER_HOST", "localhost")
	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8001"))
	if err != nil {
		logrus.Fatal(err)
	}
	conf.Server.Port = serverPort

	APIConfig = conf

	if conf.Debug {
		log.Println("Read config: ", fmt.Sprintf("%+v", conf))
	}

	return &conf, nil
}
