package http

import (
	model2 "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/mock"
	"strings"
)

type mockProxyUsecase struct {
	mock.Mock
}

func (m *mockProxyUsecase) Proxy(request *model2.Request) (response string, toRawProxy bool, err error) {
	args := m.Called(request)
	return args.String(0), args.Bool(1), args.Error(2)
}

func getStringWithoutNewLine(toAssert string) string {
	return strings.TrimSuffix(toAssert, "\n")
}
