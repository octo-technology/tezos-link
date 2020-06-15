package cache

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"log"
)

type lruBlockchainRepository struct {
	cache *lru.Cache
}

// NewLRUBlockchainRepository returns a new blockchain LRU cache repository
func NewLRUBlockchainRepository() repository.BlockchainRepository {
	cache, err := lru.New(config.ProxyConfig.Proxy.BlockchainRequestsCacheMaxItems)
	if err != nil {
		log.Fatal("could not init the LRU cache")
	}

	return &lruBlockchainRepository{
		cache: cache,
	}
}

func (l lruBlockchainRepository) Get(request *pkgmodel.Request) (interface{}, error) {
	val, ok := l.cache.Get(request.Path)
	if !ok {
		return nil, fmt.Errorf("could not get cache for path: %s", request.Path)
	}

	return val, nil
}

func (l lruBlockchainRepository) Add(request *pkgmodel.Request, response interface{}) error {
	l.cache.Add(request.Path, response)

	return nil
}
