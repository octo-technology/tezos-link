package usecases

import (
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/repository"
    "github.com/sirupsen/logrus"
)

type HealthUsecase struct{
    repo repository.ProjectRepository
}

type HealthUsecaseInterface interface{
    Health() *model.Health
}

func NewHealthUsecase(repo repository.ProjectRepository) *HealthUsecase {
    return &HealthUsecase{
        repo:repo,
    }
}

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
