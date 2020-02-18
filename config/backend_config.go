package config

import (
    "errors"
    "fmt"
    "github.com/BurntSushi/toml"
    "io/ioutil"
    "log"
)

type BackendConf struct {
    Env string
    Debug bool
    Server struct {
        Hostname string
        Port int
    }
    Db struct {
        Url string
    }
    Migration struct {
        Path string
    }
    Jwt struct {
        SignKey string
    }
}

var BackendConfig BackendConf

func ParseBackendConf(cfg string) (*BackendConf, error) {
    conf := BackendConf{}

    if data, err := ioutil.ReadFile(cfg); err != nil {
        log.Fatalf("Could not read config file:%s because of %s.", cfg, err)
    } else {
        if _, err := toml.Decode(string(data), &conf); err != nil {
            return nil, errors.New("Could not read TOML config")
        }
    }
    BackendConfig = conf
    if conf.Debug {
        log.Println("Read config: ", fmt.Sprintf("%+v", conf))
    }
    return &conf, nil
}

