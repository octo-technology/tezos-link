package cache

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
)

func TestCacheMetricsRepository_Add_Unit(t *testing.T) {
	// Given
	_, err := config.ParseProxyConf("../../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}
	cacheMetricsRepo := NewCacheMetricsRepository()

	date, err := time.Parse(time.RFC3339, "2020-03-05T10:58:56Z")
	firstMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, date)

	// When
	err = cacheMetricsRepo.Add(&firstMetric)
	nb := cacheMetricsRepo.Len()

	// Then
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, nb, 1)
}

func TestCacheMetricsRepository_GetAll_Unit(t *testing.T) {
	_, err := config.ParseProxyConf("../../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}
	// Given
	cacheMetricsRepo := NewCacheMetricsRepository()

	date, err := time.Parse(time.RFC3339, "2020-03-05T10:58:56Z")
	firstMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path",
		UUID:       "UUID",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, date)

	date, err = time.Parse(time.RFC3339, "2019-03-05T10:58:56Z")
	secondMetric := inputs.NewMetricsInput(&model.Request{
		Path:       "/random/path3",
		UUID:       "UUID2",
		Action:     0,
		RemoteAddr: "0.0.0.0",
	}, date)

	// When
	err = cacheMetricsRepo.Add(&firstMetric)
	err = cacheMetricsRepo.Add(&secondMetric)
	data, err := cacheMetricsRepo.GetAll()

	// Then
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(data), 2)
}
