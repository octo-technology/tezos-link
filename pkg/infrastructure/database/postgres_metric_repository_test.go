package database

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostgresMetricRepository_Save_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresMetricRepository(pg)
	aMetric := model.NewMetric(&model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, time.Now())

	// When
	err := pgr.Save(&aMetric)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPostgresMetricRepository_Count_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresMetricRepository(pg)
	r := &model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}
	firstMetric := model.NewMetric(r, time.Now())
	secondMetric := model.NewMetric(r, time.Now().Add(time.Duration(1)*time.Second))

	err := pgr.Save(&firstMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&secondMetric)
	if err != nil {
		t.Fatal(err)
	}

	// When
	n, err := pgr.Count("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, 2, n, "Bad count")
}
