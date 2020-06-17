package cache

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	pkginfradbinputs "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/sirupsen/logrus"
	"log"
)

type cacheMetricsRepository struct {
	cache *lru.Cache
}

// NewCacheMetricsRepository returns a new metrics LRU cache repository
func NewCacheMetricsRepository() repository.MetricInputRepository {
	configSize := config.ProxyConfig.Proxy.CacheMaxMetricItems
	cache, err := lru.New(configSize + 1)
	if err != nil {
		log.Fatal("could not init the LRU cache")
	}

	return &cacheMetricsRepository{
		cache: cache,
	}
}

func (l cacheMetricsRepository) Add(metric *pkginfradbinputs.MetricsInput) error {
	newid := l.cache.Len() + 1
	l.cache.Add(newid, metric)

	return nil
}

func (l cacheMetricsRepository) Len() int {
	return l.cache.Len()
}

func (l cacheMetricsRepository) GetAll() ([]*pkginfradbinputs.MetricsInput, error) {
	forDatabase := make([]*pkginfradbinputs.MetricsInput, 0)
	for _, key := range l.cache.Keys() {
		val, ok := l.cache.Get(key)
		if !ok {
			logrus.Debug("unable to key the metrics")
		} else {
			forDatabase = append(forDatabase, val.(*pkginfradbinputs.MetricsInput))
			l.cache.Remove(key)
		}
	}
	return forDatabase, nil
}
