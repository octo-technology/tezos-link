package database

import (
	"testing"
	"time"

	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/assert"
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
	}, time.Now().UTC())

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
	firstMetric := inputs.NewMetricsInput(r, time.Now().UTC())
	secondMetric := inputs.NewMetricsInput(r, time.Now().UTC().Add(time.Duration(1)*time.Second))

	// When there is no row
	n, err := pgr.CountAll("UUID")
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
	n, err = pgr.CountAll("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, 2, n, "Bad count")
}

func TestPostgresMetricsRepository_FindRequestsByDay_Unit(t *testing.T) {
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
	marchFixedTime, err := time.Parse(time.RFC3339, "2020-03-05T10:58:56Z")
	if err != nil {
		t.Fatal(err)
	}
	aprilFixedTime, err := time.Parse(time.RFC3339, "2020-04-05T10:58:56Z")
	if err != nil {
		t.Fatal(err)
	}
	firstMetric := inputs.NewMetricsInput(r, marchFixedTime)
	secondMetric := inputs.NewMetricsInput(r, marchFixedTime.Add(time.Duration(66)*time.Second))
	thirdMetric := inputs.NewMetricsInput(r, aprilFixedTime)

	err = pgr.Save(&firstMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&secondMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&thirdMetric)
	if err != nil {
		t.Fatal(err)
	}

	// When
	results, err := pgr.FindRequestsByDay("UUID", marchFixedTime, aprilFixedTime)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	expectedMarchRep := model.NewRequestsByDayMetrics("2020", "3", "5", 2)
	expectedAprilRep := model.NewRequestsByDayMetrics("2020", "4", "5", 1)
	expectedReps := [2]*model.RequestsByDayMetrics{expectedMarchRep, expectedAprilRep}

	assert.Equal(t, len(expectedReps), len(results))
	for i := 0; i < 1; i++ {
		assert.Equal(t, expectedReps[i], results[i])
	}
}

func TestPostgresMetricsRepository_CountRPCPaths_Unit(t *testing.T) {
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
	marchFixedTime, err := time.Parse(time.RFC3339, "2020-03-05T10:58:56Z")
	if err != nil {
		t.Fatal(err)
	}
	aprilFixedTime, err := time.Parse(time.RFC3339, "2020-04-05T10:58:56Z")
	if err != nil {
		t.Fatal(err)
	}
	firstMetric := inputs.NewMetricsInput(r, marchFixedTime)
	secondMetric := inputs.NewMetricsInput(r, marchFixedTime.Add(time.Duration(66)*time.Second))
	thirdMetric := inputs.NewMetricsInput(r, aprilFixedTime)

	err = pgr.Save(&firstMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&secondMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&thirdMetric)
	if err != nil {
		t.Fatal(err)
	}

	// When
	results, err := pgr.CountRPCPathUsage("UUID", marchFixedTime, aprilFixedTime)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	expectedRep := []*model.RPCUsageMetrics{{
		Path:  "/random/path",
		Value: 3,
	}}
	assert.Equal(t, len(expectedRep), len(results))
	assert.Equal(t, expectedRep[0], results[0])
}

func TestPostgresMetricsRepository_FindLastRequests_Unit(t *testing.T) {
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
	firstMetric := inputs.NewMetricsInput(r, time.Now().UTC())
	secondMetric := inputs.NewMetricsInput(r, time.Now().UTC().Add(time.Duration(1)*time.Second))

	// When there is no row
	n, err := pgr.FindLastRequests("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// then
	assert.Equal(t, 0, len(n), "Bad count")

	err = pgr.Save(&firstMetric)
	if err != nil {
		t.Fatal(err)
	}
	err = pgr.Save(&secondMetric)
	if err != nil {
		t.Fatal(err)
	}

	// When there is a row
	n, err = pgr.FindLastRequests("UUID")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, "/random/path", n[0], "Bad path")
	assert.Equal(t, "/random/path", n[1], "Bad path")
}

func TestPostgresMetricsRepository_SaveMany_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	date, err := time.Parse(time.RFC3339, "2020-03-05T10:58:56Z")

	pgr := NewPostgresMetricsRepository(pg)
	firstMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, date)

	date, err = time.Parse(time.RFC3339, "2019-03-05T10:58:56Z")
	secondMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path3",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, date)

	thirdMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path2",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, time.Now().UTC())

	// When
	err = pgr.SaveMany([]*inputs.MetricsInput{&firstMetric, &secondMetric, &thirdMetric})

	// Then
	if err != nil {
		t.Fatal(err)
	}

	count, err := pgr.CountAll("UUID")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, count, 3)

}
