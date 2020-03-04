package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"log"
)

func GetPostgresClient(pool dockertest.Pool) (*sql.DB, *dockertest.Resource) {
	var con *sql.DB
	var err error

	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_PASSWORD=pass", "POSTGRES_DB=tezoslink"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the postgres in the container might not be ready to accept connections yet
	url := fmt.Sprintf("postgres://postgres:pass@localhost:%s/tezoslink?sslmode=disable", resource.GetPort("5432/tcp"))
	if err := pool.Retry(func() error {
		var err error
		con, err := sql.Open("postgres", url)
		if err != nil {
			return err
		}

		return con.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	con, _ = sql.Open("postgres", url)
	runMigrations(url)

	return con, resource
}

func getDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return pool
}

func runMigrations(url string) {
	m, err := migrate.New(
		"file:../../../data/api/migrations",
		url)
	if err != nil {
		log.Fatal("Could not apply db migration: ", err)
	}
	if err := m.Up(); err != nil {
		log.Fatal("Could not apply db migration: ", err)
	}
}
