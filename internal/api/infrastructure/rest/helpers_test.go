package rest

import (
	"github.com/go-chi/chi"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockProjectUsecase struct {
	mock.Mock
}

func (m *mockProjectUsecase) SaveProject(name string) (*model.Project, error) {
	args := m.Called(name)
	return nil, args.Error(0)
}

func (m *mockProjectUsecase) FindProject(id int64) (*model.Project, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Project), nil
}

type mockHealthUsecase struct {
	mock.Mock
}

func (m *mockHealthUsecase) Health() *model.Health {
	args := m.Called()
	return args.Get(0).(*model.Health)
}

func withRouter(router *chi.Mux, f func(t *testing.T, router *chi.Mux)) func(t *testing.T) {
	return func(t *testing.T) {
		f(t, router)
	}
}

func executeRequest(req *http.Request, handler *chi.Mux) *httptest.ResponseRecorder {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	return rr
}

func buildControllerWithProjectUseCaseError(error error, ucMethod string) *Controller {
	mockHealthUsecase := &mockHealthUsecase{}
	mockProjectUsecase := &mockProjectUsecase{}
	mockProjectUsecase.
		On(ucMethod, mock.Anything).
		Return(error).
		Twice()
	rcc := NewRestController(chi.NewRouter(), mockProjectUsecase, mockHealthUsecase)
	rcc.Initialize()

	return rcc
}

func buildControllerWithProjectUseCaseReturning(project *model.Project, ucMethod string) *Controller {
	mockHealthUsecase := &mockHealthUsecase{}
	mockProjectUsecase := &mockProjectUsecase{}
	mockProjectUsecase.
		On(ucMethod, mock.Anything).
		Return(project).
		Twice()
	rcc := NewRestController(chi.NewRouter(), mockProjectUsecase, mockHealthUsecase)
	rcc.Initialize()

	return rcc
}

func buildControllerWitHealthUseCaseReturning(health *model.Health, ucMethod string) *Controller {
	mockProjectUsecase := &mockProjectUsecase{}
	mockHealthUsecase := &mockHealthUsecase{}
	mockHealthUsecase.
		On(ucMethod, mock.Anything).
		Return(health).
		Twice()
	rcc := NewRestController(chi.NewRouter(), mockProjectUsecase, mockHealthUsecase)
	rcc.Initialize()

	return rcc
}

func getStringWithoutNewLine(toAssert string) string {
	return strings.TrimSuffix(toAssert, "\n")
}
