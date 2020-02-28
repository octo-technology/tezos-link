package http

import (
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
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
	req, err := http.NewRequest("GET", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
	if err != nil {
		t.Fatal(err)
	}
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, nil, &http.Server{})
	httpController.Initialize()
	expectedRequest := model.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		model.GET,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &expectedRequest).
		Return(`{"data":{"dummy":response},"status":"success"}`, false, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then
	expectedResponse := `{"data":{"dummy":response},"status":"success"}`
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, expectedResponse, getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_Run_WhenThereIsProxiedRequest_Unit(t *testing.T) {
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
	httpController := NewHTTPController(mockProxyUsecase, rp, &http.Server{})
	httpController.Initialize()
	request := model.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		model.POST,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("", true, nil).
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
	httpController := NewHTTPController(mockProxyUsecase, rp, &http.Server{})
	httpController.Initialize()
	request := model.NewRequest(
		"/chains/main/head",
		"123e4567-e89b-12d3-a456-426655440000",
		model.POST,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("no response from proxy", false, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Bad status code")
	assert.Equal(t, "Internal Server Error", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}

func TestHttpController_ContainsRightPath_WhenPOSTRequest_Unit(t *testing.T) {
	// Given
	http.DefaultServeMux = new(http.ServeMux)
	req, err := http.NewRequest("POST", "/v1/123e4567-e89b-12d3-a456-426655440000/chains/main/number", nil)
	if err != nil {
		t.Fatal(err)
	}

	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/chains/main/number", r.URL.Path, "Bad path given to the proxy")
	}))
	defer backend.Close()
	reverseURL, err := url.Parse(backend.URL)
	if err != nil {
		t.Fatal(err)
	}
	rp := httputil.NewSingleHostReverseProxy(reverseURL)
	mockProxyUsecase := &mockProxyUsecase{}
	httpController := NewHTTPController(mockProxyUsecase, rp, &http.Server{})
	httpController.Initialize()
	request := model.NewRequest(
		"/chains/main/number",
		"123e4567-e89b-12d3-a456-426655440000",
		model.POST,
		"")

	// When
	mockProxyUsecase.
		On("Proxy", &request).
		Return("no response from proxy", true, nil).
		Once()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleProxying(httpController, "v1/"))
	handler.ServeHTTP(rr, req)

	// Then, the request is forwarded to the node
	assert.Equal(t, http.StatusOK, rr.Code, "Bad status code")
	assert.Equal(t, "", getStringWithoutNewLine(rr.Body.String()), "Bad body")
}
