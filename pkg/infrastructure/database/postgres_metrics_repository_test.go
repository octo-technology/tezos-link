package database

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostgresMetricsRepository_Save_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresMetricsRepository(pg)
	aMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, time.Now())

	// When
	err := pgr.Save(&aMetric)

	// Then
	if err != nil {
		t.Fatal(err)
	}
}

func TestPostgresMetricsRepository_Count_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresMetricsRepository(pg)
	r := &model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}
	firstMetric := inputs.NewMetricsInput(r, time.Now())
	secondMetric := inputs.NewMetricsInput(r, time.Now().Add(time.Duration(1)*time.Second))

	// When there is no row
	n, err := pgr.Count("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// then
	assert.Equal(t, 0, n, "Bad count")

	err = pgr.Save(&firstMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&secondMetric)
	if err != nil {
		t.Fatal(err)
	}

	// When there is a row
	n, err = pgr.Count("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, 2, n, "Bad count")
}
