package usecases

import (
	"errors"
	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProxyUsecase_Proxy_Unit(t *testing.T) {
	_, err := config.ParseProxyConf("../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}

	blockedRequest := pkgmodel.NewRequest("/dummy/path", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns blacklisted When there is a blacklisted path",
		testProxyUsecaseFunc(&blockedRequest, "call blacklisted", false, nil, nil, nil, nil))

	postRequest := pkgmodel.NewRequest("/chains/main/blocks/head", "UUID", pkgmodel.PUSH, "127.0.0.1")
	t.Run("Forward to reverse proxy When there is a PUSH request",
		testProxyUsecaseFunc(&postRequest, "", true, nil, nil, nil, nil))

	whitelistedCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns the cached response When there is a whitelisted path",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "Dummy cache response", false, nil, nil, nil, nil))

	t.Run("Returns the proxy response When there is no cached response",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "Dummy proxy response", false, nil, errors.New("no cache available"), nil, nil))

	whitelistedNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns the proxy response When the path is not cacheable",
		testProxyUsecaseFunc(&whitelistedNotCachedRequest, "", true, nil, nil, nil, nil))

	whitelistedCacheableNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns no response When there is no cache and an proxy error",
		testProxyUsecaseFunc(&whitelistedCacheableNotCachedRequest, "no response from proxy", false, errors.New("could not request to proxy: proxy error"),
			errors.New("no cache available"), errors.New("proxy error"), nil))
}

// Need docker-compose running in background
func TestProxyUsecase_Proxy_RedirectToMockServer_Integration(t *testing.T) {
	// Send a dummy request
	req, err := http.NewRequest("PUT", "http://0.0.0.0:8001/v1/123e4567-e89b-12d3-a456-426655440010/mockserver/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	var fb []byte
	fb, err = ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = r.Body.Close()

	assert.Equal(t, string(fb), `{
  "ports" : [ 1090 ],
  "version" : "5.9.0",
  "artifactId" : "mockserver-core",
  "groupId" : "org.mock-server"
}`)

	// Get the number of request done by this project UUID
	req, err = http.NewRequest("GET", "http://0.0.0.0:8000/api/v1/projects/123e4567-e89b-12d3-a456-426655440010/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	r, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	var sb []byte
	sb, err = ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = r.Body.Close()

	assert.Equal(t, string(sb), `{"data":{"uuid":"123e4567-e89b-12d3-a456-426655440010","requestsCount":1},"status":"success"}`)
}

func testProxyUsecaseFunc(
	r *pkgmodel.Request,
	expResp string,
	expToRawProxy bool,
	expErr error,
	cacheErr error,
	proxyErr error,
	metricErr error) func(t *testing.T) {
	return func(t *testing.T) {
		// Given
		mockCache := stubBlockchainRepository([]byte("Dummy cache response"), cacheErr)
		mockProxy := stubBlockchainRepository([]byte("Dummy proxy response"), proxyErr)
		mockMetricRepo := stubMetricRepository(metricErr)
		puc := NewProxyUsecase(mockCache, mockProxy, mockMetricRepo)

		// When
		resp, toRawProxy, err := puc.Proxy(r)

		// Then
		assert.Equal(t, expResp, resp)
		assert.Equal(t, expToRawProxy, toRawProxy)
		assert.Equal(t, expErr, err)
	}
}

func stubBlockchainRepository(response interface{}, error error) *mockBlockchainRepository {
	mockBlockchainRepository := &mockBlockchainRepository{}
	mockBlockchainRepository.
		On("Get", mock.Anything).
		Return(response, error).
		Once()
	mockBlockchainRepository.
		On("Add", mock.Anything, mock.Anything).
		Return(error).
		Once()
	return mockBlockchainRepository
}

func stubMetricRepository(error error) *mockMetricRepository {
	mockMetricRepository := &mockMetricRepository{}
	mockMetricRepository.
		On("Save", mock.Anything).
		Return(error).
		Once()
	return mockMetricRepository
}
