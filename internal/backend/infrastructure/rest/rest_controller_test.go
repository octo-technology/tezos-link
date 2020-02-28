package rest

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/backend/infrastructure/rest/inputs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestRestController_PostProject_Unit(t *testing.T) {
	// Given
	rc := buildControllerWithProjectUseCaseError(nil, "SaveProject")
	rcWithError := buildControllerWithProjectUseCaseError(errors.New("error from the DB"), "SaveProject")

	jsonBody, _ := json.Marshal(inputs.NewProject{
		Name: "New Project",
	})

	// Then
	t.Run("With expected JSON body", withRouter(rc.router,
		testPostProjectFunc(string(jsonBody), http.StatusCreated, `{"data":null,"status":"success"}`)))
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

func TestRestController_GetProjects_Unit(t *testing.T) {
	// Given
	rc := buildControllerWithProjectUseCaseError(nil, "FindProjects")
	rcWithError := buildControllerWithProjectUseCaseError(errors.New("error from the DB"), "FindProjects")

	// Then
	t.Run("Ok", withRouter(rc.router,
		testGetProjectsFunc(http.StatusOK, `{"data":null,"status":"success"}`)))
	t.Run("With a use case error", withRouter(rcWithError.router,
		testGetProjectsFunc(http.StatusBadRequest, `{"data":"error from the DB","status":"fail"}`)))
}

func testGetProjectsFunc(expectedStatus int, expectedResponse string) func(t *testing.T, router *chi.Mux) {
	return func(t *testing.T, router *chi.Mux) {
		request, err := http.NewRequest("GET", "/api/v1/projects", nil)
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
	p := model.NewProject(123, "A Project", "A_KEY")
	rc := buildControllerWithProjectUseCaseReturning(&p, "FindProject")

	// When
	request, err := http.NewRequest("GET", "/api/v1/projects/"+strconv.Itoa(int(p.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}
	requestResponse := executeRequest(request, rc.router)

	// Then
	jsonBody, err := json.Marshal(p)
	assert.Equal(t, http.StatusOK, requestResponse.Code, "Bad status code")
	assert.Equal(t, `{"data":`+string(jsonBody)+`,"status":"success"}`, getStringWithoutNewLine(requestResponse.Body.String()), "Bad body")
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
