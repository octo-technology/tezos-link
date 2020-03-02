package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/database"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest"
	"github.com/octo-technology/tezos-link/backend/internal/api/usecases"
	pkgdatabase "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database"
	"log"
	"strings"
)

var configPath = flag.String("conf", "", "Path to TOML config")

// Always run
func init() {
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Program argument --conf is required")
	} else {
		_, err := config.ParseAPIConf(*configPath)
		if err != nil {
			log.Fatalf("Could not load config from %s. Reason: %s", *configPath, err)
		}
	}

	database.Configure()
}

func main() {
	r := chi.NewRouter()
	runMigrations()

	// Repositories
	pg := database.NewPostgresProjectRepository(database.Connection)
	pm := pkgdatabase.NewPostgresMetricsRepository(database.Connection)

	// Use cases
	pu := usecases.NewProjectUsecase(pg, pm)
	hu := usecases.NewHealthUsecase(pg)

	// HTTP API
	restController := rest.NewRestController(r, pu, hu)
	restController.Initialize()
	restController.Run(config.APIConfig.Server.Port)
}

func runMigrations() {
	m, err := migrate.New(
		config.APIConfig.Migration.Path,
		config.APIConfig.Db.Url)
	if err != nil {
		log.Fatal("Could not apply db migration: ", err)
	}
	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatal("Could not apply db migration: ", err)
		}
	}
}
