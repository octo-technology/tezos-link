package usecases

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/mock"
	"time"
)

type mockProjectRepository struct {
	mock.Mock
}

func (mp *mockProjectRepository) FindByUUID(uuid string) (*model.Project, error) {
	args := mp.Called(uuid)
	return args.Get(0).(*model.Project), args.Error(1)
}

func (mp *mockProjectRepository) Save(title string, key string) (*model.Project, error) {
	args := mp.Called(title, key)
	return args.Get(0).(*model.Project), args.Error(1)
}

func (mp *mockProjectRepository) FindAll() ([]*model.Project, error) {
	args := mp.Called()
	return args.Get(0).([]*model.Project), args.Error(1)
}

func (mp *mockProjectRepository) Ping() error {
	args := mp.Called()
	return args.Error(0)
}

type mockMetricsRepository struct {
	mock.Mock
}

func (m *mockMetricsRepository) Save(metrics *inputs.MetricsInput) error {
	args := m.Called(metrics)
	return args.Error(0)
}

func (m *mockMetricsRepository) CountAll(uuid string) (int, error) {
	args := m.Called(uuid)
	return args.Int(0), args.Error(1)
}

func (m *mockMetricsRepository) FindRequestsByDay(uuid string, from time.Time, to time.Time) ([]*pkgmodel.RequestsByDayMetrics, error) {
	args := m.Called(uuid, from, to)
	return args.Get(0).([]*pkgmodel.RequestsByDayMetrics), args.Error(1)
}
