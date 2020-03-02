package usecases

import (
	"github.com/google/uuid"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
	"github.com/sirupsen/logrus"
)

// ProjectUsecase contains the project repository
type ProjectUsecase struct {
	repo repository.ProjectRepository
}

// ProjectUsecaseInterface contains all available methods of the project use-case
type ProjectUsecaseInterface interface {
	SaveProject(name string) (*model.Project, error)
	FindProject(id string) (*model.Project, error)
}

// NewProjectUsecase returns a new project use-case
func NewProjectUsecase(repo repository.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		repo: repo,
	}
}

// SaveProject save a new project
func (pu *ProjectUsecase) SaveProject(name string) (*model.Project, error) {
	p, err := pu.repo.Save(name, uuid.New().String())
	if err != nil {
		logrus.Error("Could not save project", name)
		return nil, err
	}

	logrus.Debug("Saved project", name)
	return p, nil
}

// FindProject finds a project from the project's id
func (pu *ProjectUsecase) FindProject(uuid string) (*model.Project, error) {
	p, err := pu.repo.FindByUUID(uuid)

	if err != nil {
		logrus.Error("Could not find project", p)
		return nil, err
	}

	logrus.Debug("Found project", p)
	return p, nil
}
