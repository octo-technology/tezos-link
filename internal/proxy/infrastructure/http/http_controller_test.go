package http

import (
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestHttpController_Run_WhenThereIsCachedRequest_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("GET", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/blocks/head", nil)
	if err != nil {
		t.Fatal(err)
	}
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, nil, nil, &http.Server{})
	httpController.Initialize()
	expectedRequest := pkgmodel.NewRequest(
		"/chains/main/blocks/head",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.OBTAIN,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &expectedRequest).
		Return(`{"data":{"dummy":response},"status":"success"}`, false, model.NodeTypeUnknown, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then
	expectedResponse := `{"data":{"dummy":response},"status":"success"}`
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, expectedResponse, getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_Run_WhenThereIsArchiveNodeProxiedRequest_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
	if err != nil {
		t.Fatal(err)
	}

	reverseURL, err := url.Parse("http://localhost")
	if err != nil {
		t.Fatal(err)
	}
	arp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, arp, nil, &http.Server{})
	httpController.Initialize()
	request := pkgmodel.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.PUSH,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("", true, model.ArchiveNode, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusBadGateway, rr.Code, "Bad status code")
	assert.Equal(t, "", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_Run_WhenThereIsRollingNodeProxiedRequest_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
	if err != nil {
		t.Fatal(err)
	}

	reverseURL, err := url.Parse("http://localhost")
	if err != nil {
		t.Fatal(err)
	}
	rrp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, nil, rrp, &http.Server{})
	httpController.Initialize()
	request := pkgmodel.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.PUSH,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("", true, model.RollingNode, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusBadGateway, rr.Code, "Bad status code")
	assert.Equal(t, "", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_Returns500_WhenThereIsProxyError_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
	if err != nil {
		t.Fatal(err)
	}

	reverseURL, err := url.Parse("http://localhost")
	if err != nil {
		t.Fatal(err)
	}
	rp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, rp, nil, &http.Server{})
	httpController.Initialize()
	request := pkgmodel.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.PUSH,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("no response from proxy", false, model.NodeTypeUnknown, errors.ErrNoProxyResponse).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Bad status code")
	assert.Equal(t, "Internal Server Error", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_Returns400_WhenThereIsProjectNotFoundError_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
	if err != nil {
		t.Fatal(err)
	}

	reverseURL, err := url.Parse("http://localhost")
	if err != nil {
		t.Fatal(err)
	}
	rp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, rp, nil, &http.Server{})
	httpController.Initialize()
	request := pkgmodel.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.PUSH,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("project not found", false, model.NodeTypeUnknown, errors.ErrProjectNotFound).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusNotFound, rr.Code, "Bad status code")
	assert.Equal(t, "Not Found", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_ContainsRightPath_WhenPOSTRequest_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/injection/operation", nil)
	if err != nil {
		t.Fatal(err)
	}

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/injection/operation", r.URL.Path, "Bad path given to the proxy")
	}))
	defer backend.Close()
	reverseURL, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}
	rrp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, nil, rrp, &http.Server{})
	httpController.Initialize()
	request := pkgmodel.NewRequest(
		"/injection/operation",
		"123e4567-e89b-12d3-a456-426655440000",
		pkgmodel.PUSH,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("no response from proxy", true, model.RollingNode, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, "", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_GetHealth_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, nil, nil, &http.Server{})
	httpController.Initialize()

	// When
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpController.GetHealth)
	handler.ServeHTTP(rr, req)

	// Then
	expectedResponse := `{"data":null,"status":"success"}`
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, expectedResponse, getStringWithoutNewLine(rr.Body.String()), "Bad body")
}
