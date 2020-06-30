package usecases

import (
	"time"

	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/mock"
)

type mockBlockchainRepository struct {
	mock.Mock
}

func (m *mockBlockchainRepository) Get(request *pkgmodel.Request, url string) (interface{}, error) {
	args := m.Called(request, url)
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

func (m *mockMetricsRepository) SaveMany(metrics []*inputs.MetricsInput) error {
	args := m.Called(metrics)
	return args.Error(0)
}

func (m *mockMetricsRepository) CountAll(uuid string) (int, error) {
	args := m.Called(uuid)
	return args.Int(0), args.Error(1)
}

func (m *mockMetricsRepository) Remove3MonthOldMetrics() (error) {
	args := m.Called()
	return args.Error(0)
}

func (m *mockMetricsRepository) FindRequestsByDay(uuid string, from time.Time, to time.Time) ([]*pkgmodel.RequestsByDayMetrics, error) {
	args := m.Called(uuid, from, to)
	return args.Get(0).([]*pkgmodel.RequestsByDayMetrics), args.Error(1)
}

func (m *mockMetricsRepository) CountRPCPathUsage(uuid string, from time.Time, to time.Time) ([]*pkgmodel.RPCUsageMetrics, error) {
	args := m.Called(uuid, from, to)
	return args.Get(0).([]*pkgmodel.RPCUsageMetrics), args.Error(1)
}

func (m *mockMetricsRepository) FindLastRequests(uuid string) ([]string, error) {
	args := m.Called(uuid)
	return args.Get(0).([]string), args.Error(1)
}

type mockProjectRepository struct {
	mock.Mock
}

func (mp *mockProjectRepository) FindByUUID(uuid string) (*pkgmodel.Project, error) {
	args := mp.Called(uuid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pkgmodel.Project), args.Error(1)
}

func (mp *mockProjectRepository) Save(title string, uuid string, creationDate time.Time) (*pkgmodel.Project, error) {
	args := mp.Called(title, uuid, creationDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pkgmodel.Project), args.Error(1)
}

func (mp *mockProjectRepository) FindAll() ([]*pkgmodel.Project, error) {
	args := mp.Called()
	return args.Get(0).([]*pkgmodel.Project), args.Error(1)
}

func (mp *mockProjectRepository) Ping() error {
	args := mp.Called()
	return args.Error(0)
}

type mockCacheMetricsRepository struct {
	mock.Mock
}

func (m *mockCacheMetricsRepository) Add(metrics *inputs.MetricsInput) error {
	args := m.Called(metrics)
	return args.Error(0)
}

func (m *mockCacheMetricsRepository) GetAll() ([]*inputs.MetricsInput, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*inputs.MetricsInput), args.Error(1)
}

func (m *mockCacheMetricsRepository) Len() int {
	args := m.Called()
	return args.Get(0).(int)
}
