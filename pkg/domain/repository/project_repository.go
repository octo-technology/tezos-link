package repository

import (
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"time"
)

// ProjectRepository contains all methods available for the project repository
type ProjectRepository interface {
	FindAll() ([]*pkgmodel.Project, error)
	FindByUUID(uuid string) (*pkgmodel.Project, error)
	Save(title string, uuid string, creationDate time.Time, network string) (*pkgmodel.Project, error)
	Ping() error
}
