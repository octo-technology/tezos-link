package usecases

import (
	"errors"
	"github.com/google/uuid"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgrepository "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/sirupsen/logrus"
)

// ProjectUsecase contains the project repository
type ProjectUsecase struct {
	projectRepo repository.ProjectRepository
	metricsRepo pkgrepository.MetricsRepository
}

// ProjectUsecaseInterface contains all available methods of the project use-case
type ProjectUsecaseInterface interface {
	CreateProject(name string) (*model.Project, error)
	FindProjectAndMetrics(uuid string) (*model.Project, *model.Metrics, error)
}

// NewProjectUsecase returns a new project use-case
func NewProjectUsecase(rp repository.ProjectRepository, mr pkgrepository.MetricsRepository) *ProjectUsecase {
	return &ProjectUsecase{
		projectRepo: rp,
		metricsRepo: mr,
	}
}

// CreateProject create and save a new project
func (pu *ProjectUsecase) CreateProject(name string) (*model.Project, error) {
	p, err := pu.projectRepo.Save(name, uuid.New().String())
	if name == "" {
		logrus.Error("empty project name", name)
		return nil, modelerrors.ErrNoProjectName
	}

	// TODO Add the creation date
	p, err = pu.projectRepo.Save(name, uuid.New().String())
	if err != nil {
		logrus.Error("Could not save project", name)
		return nil, err
	}

	logrus.Debug("Saved project", name)
	return p, nil
}

// FindProjectAndMetrics finds a project and the associated metrics of a given project's uuid
func (pu *ProjectUsecase) FindProjectAndMetrics(uuid string) (*model.Project, *model.Metrics, error) {
	p, err := pu.projectRepo.FindByUUID(uuid)

	if errors.Is(err, modelerrors.ErrProjectNotFound) {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	n, err := pu.metricsRepo.Count(uuid)
	if err != nil {
		return nil, nil, err
	}

	m := model.NewMetrics(n)
	logrus.Debug("Found project and metrics", p, m)
	return p, &m, nil
}
