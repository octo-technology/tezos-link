package repository

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"time"
)

// ProjectRepository contains all methods available for the project repository
type ProjectRepository interface {
	FindAll() ([]*model.Project, error)
	FindByUUID(uuid string) (*model.Project, error)
	Save(title string, uuid string, creationDate time.Time) (*model.Project, error)
	Ping() error
}
