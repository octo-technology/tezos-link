package repository

import "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"

// ProjectRepository contains all methods available for the project repository
type ProjectRepository interface {
	FindAll() ([]*model.Project, error)
	FindByID(id int64) (*model.Project, error)
	Save(name string, key string) (*model.Project, error)
	UpdateByID(project *model.Project) error
	DeleteByID(project *model.Project) error
	Ping() error
}
