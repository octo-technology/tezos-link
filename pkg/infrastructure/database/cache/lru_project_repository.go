package cache

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/octo-technology/tezos-link/backend/config"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"log"
	"time"
)

type lruProjectRepository struct {
	cache *lru.Cache
}

// NewLruProjectRepository returns a new project LRU cache repository
func NewLruProjectRepository() repository.ProjectRepository {
	cache, err := lru.New(config.ProxyConfig.Proxy.ProjectsCacheMaxItems)
	if err != nil {
		log.Fatal("could not init the LRU cache")
	}

	return &lruProjectRepository{
		cache: cache,
	}
}

func (l lruProjectRepository) FindAll() ([]*pkgmodel.Project, error) {
	panic("implement me")
}

func (l lruProjectRepository) FindByUUID(uuid string) (*pkgmodel.Project, error) {
	val, ok := l.cache.Get(uuid)
	if !ok {
		return nil, fmt.Errorf("could not get cache for path: %s", uuid)
	}

	return val.(*pkgmodel.Project), nil
}

func (l lruProjectRepository) Save(title string, uuid string, creationDate time.Time) (*pkgmodel.Project, error) {
	noID := 0
	project := pkgmodel.NewProject(int64(noID), title, uuid, creationDate)

	l.cache.Add(uuid, project)

	return &project, nil
}

func (l lruProjectRepository) Ping() error {
	panic("implement me")
}
