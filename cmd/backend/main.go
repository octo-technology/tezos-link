package main

import (
    "flag"
    "github.com/go-chi/chi"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    _ "github.com/lib/pq"
    "github.com/octo-technology/tezos-link/backend/config"
    "github.com/octo-technology/tezos-link/backend/internal/backend/infrastructure/database"
    "github.com/octo-technology/tezos-link/backend/internal/backend/infrastructure/rest"
    "github.com/octo-technology/tezos-link/backend/internal/backend/usecases"
    "log"
)

var configPath = flag.String("conf", "", "Path to TOML config")

// Always run
func init() {
    flag.Parse()

    if *configPath == "" {
        log.Fatal("Program argument --conf is required")
    } else {
        _, err := config.Parse(*configPath)
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

    // Use cases
    pu := usecases.NewProjectUsecase(pg)
    hu := usecases.NewHealthUsecase(pg)

    // HTTP API
    restController := rest.NewRestController(r, pu, hu)
    restController.Initialize()
    restController.Run(config.Conf.Server.Port)
}

func runMigrations() {
    m, err := migrate.New(
        config.Conf.Migration.Path,
        config.Conf.Db.Url)
    if err != nil {
        log.Fatal("Could not apply db migration: ", err)
    }
    if err := m.Up(); err != nil {
        log.Fatal("Could not apply db migration: ", err)
    }
}
