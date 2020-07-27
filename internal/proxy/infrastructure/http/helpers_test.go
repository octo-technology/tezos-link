package http

import (
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/mock"
	"strings"
)

type mockProxyUsecase struct {
	mock.Mock
}

func (m *mockProxyUsecase) Proxy(request *pkgmodel.Request) (response string, toRawProxy bool, nodeType model.NodeType, err error) {
	args := m.Called(request)
	return args.String(0), args.Bool(1), args.Get(2).(model.NodeType), args.Error(3)
}

func getStringWithoutNewLine(toAssert string) string {
	return strings.TrimSuffix(toAssert, "\n")
}
