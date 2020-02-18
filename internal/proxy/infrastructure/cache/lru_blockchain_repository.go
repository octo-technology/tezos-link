package cache

import (
    "errors"
    "fmt"
    lru "github.com/hashicorp/golang-lru"
    "github.com/octo-technology/tezos-link/backend/config"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
    "log"
)

type lruBlockchainRepository struct {
    cache *lru.Cache
}

func NewLruBlockchainRepository() repository.BlockchainRepository {
    cache, err := lru.New(config.ProxyConfig.Proxy.CacheMaxItems)
    if err != nil {
      log.Fatal("could not init the LRU cache")
    }

    return &lruBlockchainRepository{
        cache: cache,
    }
}

func (l lruBlockchainRepository) Get(request *model.Request) (interface{}, error) {
    val, ok := l.cache.Get(request.Path)
    if ok {
        return nil, errors.New(fmt.Sprintf("could not get cache for path: %s", request.Path))
    }

    return val, nil
}

func (l lruBlockchainRepository) Add(request *model.Request, response interface{}) error {
    l.cache.Add(request.Path, response)

    return nil
}
