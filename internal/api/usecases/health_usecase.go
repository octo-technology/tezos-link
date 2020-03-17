package usecases

import (
    "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
    repository2 "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
    "github.com/sirupsen/logrus"
)

// HealthUsecase contains the project repository to do the health check
type HealthUsecase struct {
	projectRepo repository2.ProjectRepository
}

// HealthUsecaseInterface contains all available methods of the health use-case
type HealthUsecaseInterface interface {
	Health() *model.Health
}

// NewHealthUsecase returns a new health use-case
func NewHealthUsecase(repo repository2.ProjectRepository) *HealthUsecase {
	return &HealthUsecase{
		projectRepo: repo,
	}
}

// Health checks and returns true if the app can ping the database
func (hu *HealthUsecase) Health() *model.Health {
	err := hu.projectRepo.Ping()

	if err != nil {
		logrus.Error("Could not ping DB", err)
		health := model.NewHealth(false)
		return &health
	}

	logrus.Debug("Successfully pinged DB")
	health := model.NewHealth(true)
	return &health
}
