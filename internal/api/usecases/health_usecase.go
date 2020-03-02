package usecases

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
	"github.com/sirupsen/logrus"
)

// HealthUsecase contains the project repository to do the health check
type HealthUsecase struct {
	repo repository.ProjectRepository
}

// HealthUsecaseInterface contains all available methods of the health use-case
type HealthUsecaseInterface interface {
	Health() *model.Health
}

// NewHealthUsecase returns a new health use-case
func NewHealthUsecase(repo repository.ProjectRepository) *HealthUsecase {
	return &HealthUsecase{
		repo: repo,
	}
}

// Health checks and returns true if the app can ping the database
func (hu *HealthUsecase) Health() *model.Health {
	err := hu.repo.Ping()

	if err != nil {
		logrus.Error("Could not ping DB", err)
		health := model.NewHealth(false)
		return &health
	}

	logrus.Debug("Successfully pinged DB")
	health := model.NewHealth(true)
	return &health
}
