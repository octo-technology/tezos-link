package usecases

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/sirupsen/logrus"
)

// MetricUsecase contains the project repository to do the health check
type MetricUsecase struct {
	repo repository.MetricRepository
}

// MetricUsecaseInterface contains all available methods of the metric use-case
type MetricUsecaseInterface interface {
	CountRequests(uuid string) (int, error)
}

// NewMetricUsecase returns a new metric use-case
func NewMetricUsecase(repo repository.MetricRepository) *MetricUsecase {
	return &MetricUsecase{
		repo: repo,
	}
}

// CountRequests count all requests done by a project ID
func (mu *MetricUsecase) CountRequests(uuid string) (int, error) {
	n, err := mu.repo.Count(uuid)

	if err != nil {
		logrus.Error(err)
	}

	return n, err
}
