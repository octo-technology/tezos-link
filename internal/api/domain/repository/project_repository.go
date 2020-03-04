package repository

import "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"

// ProjectRepository contains all methods available for the project repository
type ProjectRepository interface {
	FindAll() ([]*model.Project, error)
	FindByUUID(uuid string) (*model.Project, error)
	Save(title string, key string) (*model.Project, error)
	UpdateByID(project *model.Project) error
	DeleteByID(project *model.Project) error
	Ping() error
}
