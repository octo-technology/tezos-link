package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
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
func (pg postgresMetricsRepository) Save(metricInput *inputs.MetricsInput) error {
	_, err := pg.connection.
		Exec("INSERT INTO metrics(path, uuid, remote_address, date_request) VALUES ($1, $2, $3, $4)",
			metricInput.Request.Path,
			metricInput.Request.UUID,
			metricInput.Request.RemoteAddr,
			metricInput.Date)

	if err != nil {
		return fmt.Errorf("could not insert metricInput for UUID %s: %s", metricInput.Request.UUID, err)
	}

	return nil
}

// insert many metrics
func (pg postgresMetricsRepository) SaveMany(metricInputs []*inputs.MetricsInput) error {
	txn, err := pg.connection.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn("metrics", "path", "uuid", "remote_address", "date_request"))
	if err != nil {
		return err
	}

	for _, metricInput := range metricInputs {
		_, err = stmt.Exec(metricInput.Request.Path, metricInput.Request.UUID, metricInput.Request.RemoteAddr, metricInput.Date)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec() // to flush data
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
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

func (pg postgresMetricsRepository) Remove3MonthOldMetrics() error {
	_, err := pg.connection.
		Exec("DELETE FROM metrics WHERE date_request <= now() - INTERVAL '3 MONTH'")

	if errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("could not get 3-month old requests : %s", err)
		return modelerrors.ErrNoMetricsFound
	}

	if err != nil {
		return err
	}

	return nil
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
		"path, "+
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

func (pg postgresMetricsRepository) FindLastRequests(uuid string) ([]string, error) {
	rows, err := pg.connection.Query("SELECT "+
		"path "+
		"FROM metrics "+
		"WHERE (uuid = $1) "+
		"ORDER BY date_request DESC "+
		"LIMIT 10", uuid)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve metrics: %s", err)
	}

	var r []string
	for rows.Next() {
		var path string
		err := rows.Scan(&path)
		if err != nil {
			return nil, fmt.Errorf("could not map metrics: %s", err)
		}
		r = append(r, path)
	}

	return r, nil
}
