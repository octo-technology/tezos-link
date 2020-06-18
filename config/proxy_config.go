package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type ProxyConf struct {
	Debug bool
	Tezos struct {
		ArchiveHost        string
		Port               int
		RollingHost        string
		RollingPort        int
		WhitelistedRolling []string
	}
	Server struct {
		Port int
	}
	Database struct {
		Url string
	}
	Proxy struct {
		ReadTimeout                     int
		WriteTimeout                    int
		IdleTimeout                     int
		WhitelistedMethods              []string
		BlockedMethods                  []string
		DontCache                       []string
		RateLimitPeriod                 int
		RateLimitCount                  int64
		BlockchainRequestsCacheMaxItems int
		ProjectsCacheMaxItems           int
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

	dbUrl := getEnv("DATABASE_URL", "postgres:5432")
	dbUser := getEnv("DATABASE_USERNAME", "user")
	dbPass := getEnv("DATABASE_PASSWORD", "pass")
	dbTable := getEnv("DATABASE_TABLE", "tezoslink?sslmode=disable")
	dbParam := getEnv("DATABASE_ADDITIONAL_PARAMETER", "sslmode=disable")
	conf.Database.Url = fmt.Sprintf("postgres://%s:%s@%s/%s?%s", dbUser, dbPass, dbUrl, dbTable, dbParam)

	conf.Tezos.ArchiveHost = getEnv("ARCHIVE_NODES_URL", "node")
	conf.Tezos.RollingHost = getEnv("ROLLING_NODES_URL", "node")
	tezosPort, err := strconv.Atoi(getEnv("TEZOS_PORT", "1090"))
	if err != nil {
		logrus.Fatal(err)
	}
	tezosRollingPort, err := strconv.Atoi(getEnv("TEZOS_ROLLING_PORT", "1090"))
	if err != nil {
		logrus.Fatal(err)
	}
	conf.Tezos.Port = tezosPort
	conf.Tezos.RollingPort = tezosRollingPort

	serverPort, err := strconv.Atoi(getEnv("SERVER_PORT", "8001"))
	if err != nil {
		logrus.Fatal(err)
	}
	conf.Server.Port = serverPort

	ProxyConfig = conf
	if conf.Debug {
		log.Println("Read config: ", fmt.Sprintf("%+v", conf))
	}

	return &conf, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
