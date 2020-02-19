package http

import (
    "fmt"
    "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
    "github.com/stretchr/testify/assert"
    "log"
    "net/http"
    "net/http/httptest"
    "net/http/httputil"
    "net/url"
    "testing"
)

func TestHttpController_Run_WhenThereIsCachedRequest(t *testing.T) {
    // Given
    http.DefaultServeMux = new(http.ServeMux)
    req, err := http.NewRequest("GET", "v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
    if err != nil {
        t.Fatal(err)
    }
    mockProxyUsecase := &mockProxyUsecase{}
    httpController := NewHttpController(mockProxyUsecase, nil, &http.Server{})
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

func TestHttpController_Run_WhenThereIsProxiedRequest(t *testing.T) {
    // Given
    http.DefaultServeMux = new(http.ServeMux)
    req, err := http.NewRequest("POST", "v1/123e4567-e89b-12d3-a456-426655440000/chains/main/head", nil)
    if err != nil {
        t.Fatal(err)
    }

    reverseUrl, err := url.Parse("http://localhost")
    if err != nil {
        log.Fatal(fmt.Sprintf("could not read blockchain node reverse url from configuration: %s", err))
    }
    rp := httputil.NewSingleHostReverseProxy(reverseUrl)
    mockProxyUsecase := &mockProxyUsecase{}
    httpController := NewHttpController(mockProxyUsecase, rp, &http.Server{})
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
