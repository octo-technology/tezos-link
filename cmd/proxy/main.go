package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/cache"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/database"
	httpinfra "github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/http"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/infrastructure/proxy"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/usecases"
	pkgcache "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/cache"
	pkgdatabase "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database"
	"github.com/sirupsen/logrus"
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

	database.Configure()
}

func writeCachedRequestsRoutine(p *usecases.ProxyUsecase) {
	for true {
		p.WriteCachedRequestsRoutine()
	}
}

func main() {
	reverseURL, err := url.Parse("http://" + config.ProxyConfig.Tezos.Host + ":" + strconv.Itoa(config.ProxyConfig.Tezos.Port))
	if err != nil {
		log.Fatal(fmt.Sprintf("could not read blockchain node reverse url from configuration: %s", err))
	}
	logrus.Info("proxying requests to node: ", reverseURL)
	reverseProxy := httputil.NewSingleHostReverseProxy(reverseURL)

	// Repositories
	cacheBlockchainRepo := cache.NewCacheBlockchainRepository()
	proxyRepo := proxy.NewProxyBlockchainRepository()
	projectRepo := pkgdatabase.NewPostgresProjectRepository(database.Connection)
	cacheProjectRepo := pkgcache.NewLRUProjectRepository()
	metricsRepo := pkgdatabase.NewPostgresMetricsRepository(database.Connection)
	cacheMetricsRepo := cache.NewCacheMetricsRepository()

	cacheMetricsRepo := cache.NewCacheMetricsRepository()

	// Use cases
<<<<<<< HEAD
	proxyUsecase := usecases.NewProxyUsecase(cacheBlockchainRepo, proxyRepo, metricsRepo, projectRepo, cacheMetricsRepo)
=======
	proxyUsecase := usecases.NewProxyUsecase(cacheBlockchainRepo, proxyRepo, metricsRepo, projectRepo, cacheProjectRepo, cacheMetricsRepo)
>>>>>>> b6bf93d452ef683c38f71adbc215e1d7dafde8a6

	// Routine
	go writeCachedRequestsRoutine(proxyUsecase)

	// HTTP API
	server := http.Server{
		Addr:         ":" + strconv.Itoa(config.ProxyConfig.Server.Port),
		ReadTimeout:  time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ProxyConfig.Proxy.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.ProxyConfig.Proxy.IdleTimeout) * time.Second,
	}
	httpController := httpinfra.NewHTTPController(proxyUsecase, reverseProxy, &server)
	httpController.Initialize()
	httpController.Run()
}
