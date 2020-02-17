package repository

import "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"

type ProjectRepository interface {
    FindAll() ([]*model.Project, error)
    FindById(id int64) (*model.Project, error)
    Save(name string, key string) (*model.Project, error)
    UpdateById(claim *model.Project) error
    DeleteById(claim *model.Project) error
    Ping() error
}
