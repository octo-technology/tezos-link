package database

import (
	"database/sql"
	"errors"
	"fmt"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/sirupsen/logrus"
)

type postgresMetricsRepository struct {
	connection *sql.DB
}

// NewPostgresMetricsRepository returns a new postgres metrics repository
func NewPostgresMetricsRepository(connection *sql.DB) repository.MetricsRepository {
	return &postgresMetricsRepository{
		connection: connection,
	}
}

// Save insert a new metrics
func (pg postgresMetricsRepository) Save(metrics *inputs.MetricsInput) error {
	_, err := pg.connection.
		Exec("INSERT INTO metrics(path, uuid, remote_address, date_request) VALUES ($1, $2, $3, $4)",
			metrics.Request.Path,
			metrics.Request.UUID,
			metrics.Request.RemoteAddr,
			metrics.Date)

	if err != nil {
		return fmt.Errorf("could not insert metrics for UUID %s: %s", metrics.Request.UUID, err)
	}

	return nil
}

// Count count request for a specified project uuid
func (pg postgresMetricsRepository) Count(uuid string) (int, error) {
	var count int

	err := pg.connection.
		QueryRow("SELECT COUNT(*) FROM metrics WHERE uuid = $1", uuid).
		Scan(&count)

	if errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("could not count requests for UUID %s: %s", uuid, err)
		return -1, modelerrors.ErrProjectNotFound
	}

	return count, nil
}
