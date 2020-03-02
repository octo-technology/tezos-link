package database

import (
	"database/sql"
	"fmt"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
)

type postgresMetricRepository struct {
	connection *sql.DB
}

// NewPostgresMetricRepository returns a new postgres metric repository
func NewPostgresMetricRepository(connection *sql.DB) repository.MetricRepository {
	return &postgresMetricRepository{
		connection: connection,
	}
}

// Save insert a new metric
func (pg postgresMetricRepository) Save(metric *model.Metric) error {
	_, err := pg.connection.
		Exec("INSERT INTO metrics(path, uuid, remote_address, date_request) VALUES ($1, $2, $3, $4)",
			metric.Request.Path,
			metric.Request.UUID,
			metric.Request.RemoteAddr,
			metric.Date)

	if err != nil {
		return fmt.Errorf("could not insert metric for UUID %s: %s", metric.Request.UUID, err)
	}

	return nil
}

// Count count request for a specified uuid
func (pg postgresMetricRepository) Count(uuid string) (int, error) {
	var n int

	err := pg.connection.
		QueryRow("SELECT COUNT(*) FROM metrics WHERE uuid = $1", uuid).Scan(&n)

	if err != nil {
		return -1, fmt.Errorf("could not count requests for UUID %s: %s", uuid, err)
	}

	return n, nil
}
