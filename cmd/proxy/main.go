package main

import (
    "flag"
    "fmt"
    "github.com/octo-technology/tezos-link/backend/config"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/cache"
    http_infra "github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/http"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/proxy"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/usecases"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strconv"
    "time"
)

var configPath = flag.String("conf", "", "Path to TOML config")

// Always run
func init() {
    flag.Parse()

    if *configPath == "" {
        log.Fatal("Program argument --conf is required")
    } else {
        _, err := config.ParseProxyConf(*configPath)
        if err != nil {
            log.Fatalf("Could not load config from %s. Reason: %s", *configPath, err)
        }
    }

    //database.Configure()
}

func main() {
    reverseUrl, err := url.Parse("http://" + config.ProxyConfig.Tezos.Host + ":"+ strconv.Itoa(config.ProxyConfig.Tezos.Port))
    if err != nil {
        log.Fatal(fmt.Sprintf("could not read blockchain node reverse url from configuration: %s", err))
    }
    rp := httputil.NewSingleHostReverseProxy(reverseUrl)
    //runMigrations()

    // Repositories
    cb := cache.NewLruBlockchainRepository()
    pb := proxy.NewProxyBlockchainRepository()

    // Use cases
    pu := usecases.NewProxyUsecase(cb, pb)

    // HTTP API
    srv := http.Server{
        Addr:         ":" + strconv.Itoa(config.ProxyConfig.Server.Port),
        ReadTimeout:  time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second,
        WriteTimeout: time.Duration(config.ProxyConfig.Proxy.WriteTimeout) * time.Second,
        IdleTimeout:  time.Duration(config.ProxyConfig.Proxy.IdleTimeout) * time.Second,
    }
    httpController := http_infra.NewHttpController(pu, rp, &srv)
    httpController.Initialize()
    httpController.Run()
}

/*
func runMigrations() {
    m, err := migrate.New(
        config.BackendConfig.Migration.Path,
        config.BackendConfig.Db.Url)
    if err != nil {
        log.Fatal("Could not apply db migration: ", err)
    }
    if err := m.Up(); err != nil {
        log.Fatal("Could not apply db migration: ", err)
    }
}
 */
