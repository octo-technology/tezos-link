package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
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
	APIConfig = conf
	if conf.Debug {
		log.Println("Read config: ", fmt.Sprintf("%+v", conf))
	}
	return &conf, nil
}
