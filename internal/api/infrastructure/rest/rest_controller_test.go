package rest

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/inputs"
	modelerrors "github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRestController_PostProject_Unit(t *testing.T) {
	// Given
	creationDate := time.Now().UTC()
	p := pkgmodel.NewProject(1, "PROJECT_NAME", "AN_UUID", creationDate)
	rc := buildControllerWithProjectUseCaseError(&p, nil, "CreateProject")
	rcWithError := buildControllerWithProjectUseCaseError(nil, errors.New("error from the DB"), "CreateProject")
	rcWithEmptyNameError := buildControllerWithProjectUseCaseError(nil, modelerrors.ErrNoProjectName, "CreateProject")

	jsonBody, _ := json.Marshal(inputs.NewProject{
		Title: "New Project",
	})
	unexpectedJSONBody := `{"BADDDDD":"BAD_KEY"}`

	// Then
	t.Run("With expected JSON body", withRouter(rc.router,
		testPostProjectFunc(string(jsonBody), http.StatusCreated, `{"data":null,"status":"success"}`)))
	t.Run("With unexpected JSON body", withRouter(rc.router,
		testPostProjectFunc(unexpectedJSONBody, http.StatusBadRequest, `{"data":"json: unknown field \"BADDDDD\"","status":"fail"}`)))
	t.Run("With empty JSON body", withRouter(rcWithEmptyNameError.router,
		testPostProjectFunc("{}", http.StatusBadRequest, `{"data":"project name not defined","status":"fail"}`)))
	t.Run("With empty body", withRouter(rc.router,
		testPostProjectFunc("", http.StatusBadRequest, `{"data":"EOF","status":"fail"}`)))
	t.Run("With a use case error", withRouter(rcWithError.router,
		testPostProjectFunc(string(jsonBody), http.StatusBadRequest, `{"data":"error from the DB","status":"fail"}`)))
}

func testPostProjectFunc(jsonInput string, expectedStatus int, expectedResponse string) func(t *testing.T, router *chi.Mux) {
	return func(t *testing.T, router *chi.Mux) {
		request, err := http.NewRequest("POST", "/api/v1/projects", strings.NewReader(jsonInput))
		if err != nil {
			t.Fatal(err)
		}
		requestResponse := executeRequest(request, router)

		assert.Equal(t, expectedStatus, requestResponse.Code, "Bad status code")
		assert.Equal(t, expectedResponse, getStringWithoutNewLine(requestResponse.Body.String()), "Bad body")
	}
}

func TestRestController_GetProject_Unit(t *testing.T) {
	// Given
	firstRequestsMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "1", 4)
	secondRequestsMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "12", 5)
	rpcUsage := pkgmodel.NewRPCUsageMetrics("/dummy/path", 3)
	creationDate := time.Now().UTC()
	stubRPCUSageMetrics := []*pkgmodel.RPCUsageMetrics{rpcUsage}
	stubRequestsMetrics := []*pkgmodel.RequestsByDayMetrics{firstRequestsMetrics, secondRequestsMetrics}
	p := pkgmodel.NewProject(123, "A Project", "A_UUID_666", creationDate)
	m := pkgmodel.NewMetrics(3, stubRequestsMetrics, stubRPCUSageMetrics, []string{"/a/path", "/another/path"})

	mockProjectUsecase := &mockProjectUsecase{}
	mockProjectUsecase.
		On("FindProjectAndMetrics", "A_UUID_666", mock.Anything, mock.Anything).
		Return(&p, &m, nil).
		Once()
	mockHealthUsecase := &mockHealthUsecase{}
	rcc := NewRestController(chi.NewRouter(), mockProjectUsecase, mockHealthUsecase)
	rcc.Initialize()

	// When
	request, err := http.NewRequest("GET", "/api/v1/projects/A_UUID_666", nil)
	if err != nil {
		t.Fatal(err)
	}
	requestResponse := executeRequest(request, rcc.router)

	// Then
	assert.Equal(t, http.StatusOK, requestResponse.Code, "Bad status code")
	assert.Equal(t, `{"data":{"title":"A Project","uuid":"A_UUID_666","metrics":{"requestsCount":3,"requestsByDay":[{"date":"2014-11-1","value":4},{"date":"2014-11-12","value":5}],"rpcUsage":[{"id":"/dummy/path","label":"/dummy/path","value":3}],"lastRequests":["/a/path","/another/path"]}},"status":"success"}`, getStringWithoutNewLine(requestResponse.Body.String()), "Bad body")
}

func TestRestController_GetHealth_Unit(t *testing.T) {
	// Given
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	health := model.NewHealth(true)
	rc := buildControllerWitHealthUseCaseReturning(&health, "Health")

	// When
	rr := executeRequest(req, rc.router)

	// Then
	expectedResponse := `{"data":{"connectedToDb":true},"status":"success"}`
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, expectedResponse, getStringWithoutNewLine(rr.Body.String()), "Bad body")
}
