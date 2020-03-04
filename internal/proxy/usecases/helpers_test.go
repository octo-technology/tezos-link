package usecases

import (
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/mock"
)

type mockBlockchainRepository struct {
	mock.Mock
}

func (m *mockBlockchainRepository) Get(request *pkgmodel.Request) (interface{}, error) {
	args := m.Called(request)
	return args.Get(0), args.Error(1)
}

func (m *mockBlockchainRepository) Add(request *pkgmodel.Request, response interface{}) error {
	_ = m.Called(request, response)
	return nil
}

type mockMetricsRepository struct {
	mock.Mock
}

func (m *mockMetricsRepository) Save(metrics *inputs.MetricsInput) error {
	args := m.Called(metrics)
	return args.Error(0)
}

func (m *mockMetricsRepository) Count(uuid string) (int, error) {
	args := m.Called(uuid)
	return args.Int(0), args.Error(1)
}
