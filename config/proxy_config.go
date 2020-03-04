package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type ProxyConf struct {
	Debug bool
	Tezos struct {
		Host string
		Port int
	}
	Server struct {
		Port int
	}
	Database struct {
		Url string
	}
	Proxy struct {
		ReadTimeout        int
		WriteTimeout       int
		IdleTimeout        int
		WhitelistedMethods []string
		BlockedMethods     []string
		DontCache          []string
		RateLimitPeriod    int
		RateLimitCount     int64
		CacheMaxItems      int
	}
}

var ProxyConfig ProxyConf

func ParseProxyConf(cfg string) (*ProxyConf, error) {
	conf := ProxyConf{}

	if data, err := ioutil.ReadFile(cfg); err != nil {
		log.Fatalf("Could not read config file:%s because of %s.", cfg, err)
	} else {
		if _, err := toml.Decode(string(data), &conf); err != nil {
			return nil, errors.New("could not read TOML config")
		}
	}
	ProxyConfig = conf
	if conf.Debug {
		log.Println("Read config: ", fmt.Sprintf("%+v", conf))
	}
	return &conf, nil
}
