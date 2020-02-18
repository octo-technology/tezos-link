package usecases

import (
    "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
    "github.com/stretchr/testify/mock"
)

type mockBlockchainRepository struct {
    mock.Mock
}

func (m *mockBlockchainRepository) Get(request *model.Request) (interface{}, error) {
    args := m.Called(request)
    return args.Get(0), args.Error(1)
}

func (m *mockBlockchainRepository) Add(request *model.Request, response interface{}) error {
    _ = m.Called(request, response)
    return nil
}
