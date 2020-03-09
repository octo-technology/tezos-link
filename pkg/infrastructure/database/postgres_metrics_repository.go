package database

import (
	"database/sql"
	"errors"
	"fmt"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/sirupsen/logrus"
	"time"
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

// CountAll count all requests for a specified project uuid
func (pg postgresMetricsRepository) CountAll(uuid string) (int, error) {
	var count int

	err := pg.connection.
		QueryRow("SELECT COUNT(*) FROM metrics WHERE uuid = $1", uuid).
		Scan(&count)

	if errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("could not count requests for UUID %s: %s", uuid, err)
		return -1, modelerrors.ErrProjectNotFound
	}

	if err != nil {
		return -1, err
	}

	return count, nil
}

func (pg postgresMetricsRepository) FindRequestsByDay(uuid string, from time.Time, to time.Time) ([]*model.RequestsByDayMetrics, error) {
	rows, err := pg.connection.Query("SELECT "+
		"EXTRACT(month from date_request) AS mon, "+
		"EXTRACT(year from date_request) AS yyyy, "+
		"EXTRACT(day from date_request) AS dd, "+
		"COUNT(*) "+
		"FROM metrics "+
		"WHERE (uuid = $1) AND (date_request BETWEEN $2 AND $3) "+
		"GROUP BY 1, 2, 3", uuid, from, to)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve metrics: %s", err)
	}

	var r []*model.RequestsByDayMetrics
	for rows.Next() {
		cur := model.RequestsByDayMetrics{}
		err := rows.Scan(&cur.Month, &cur.Year, &cur.Day, &cur.Value)
		if err != nil {
			return nil, fmt.Errorf("could not map metrics: %s", err)
		}
		r = append(r, &cur)
	}

	return r, nil
}

func (pg postgresMetricsRepository) CountRPCPathUsage(uuid string, from time.Time, to time.Time) ([]*model.RPCUsageMetrics, error) {
	rows, err := pg.connection.Query("SELECT "+
		"path, " +
		"COUNT(*) "+
		"FROM metrics "+
		"WHERE (uuid = $1) AND (date_request BETWEEN $2 AND $3) "+
		"GROUP BY 1", uuid, from, to)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve metrics: %s", err)
	}

	var r []*model.RPCUsageMetrics
	for rows.Next() {
		cur := model.RPCUsageMetrics{}
		err := rows.Scan(&cur.Path, &cur.Value)
		if err != nil {
			return nil, fmt.Errorf("could not map metrics: %s", err)
		}
		r = append(r, &cur)
	}

	return r, nil
}
