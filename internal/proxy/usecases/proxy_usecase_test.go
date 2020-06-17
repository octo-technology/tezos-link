package usecases

import (
	"encoding/json"
	"errors"
	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/inputs"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/outputs"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	dbinputs "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestProxyUsecase_Proxy_Unit(t *testing.T) {
	_, err := config.ParseProxyConf("../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}

	blockedRequest := pkgmodel.NewRequest("/dummy/path", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns blacklisted When there is a blacklisted path",
		testProxyUsecaseFunc(&blockedRequest, "call blacklisted", false, nil, nil, nil, nil, nil))

	postRequest := pkgmodel.NewRequest("/chains/main/blocks/head", "UUID", pkgmodel.PUSH, "127.0.0.1")
	t.Run("Forward to reverse proxy When there is a PUSH request",
		testProxyUsecaseFunc(&postRequest, "", true, nil, nil, nil, nil, nil))

	whitelistedCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns the cached response When there is a whitelisted path",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "Dummy cache response", false, nil, nil, nil, nil, nil))

	t.Run("Returns the proxy response When there is no cached response",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "Dummy proxy response", false, nil, errors.New("no cache available"), nil, nil, nil))

	t.Run("Returns the proxy response When there is no cached response",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "Dummy proxy response", false, nil, errors.New("no cache available"), nil, nil, nil))

	projectNotFound := errors.New("project not found")
	t.Run("Returns no project found error When there is a bad project UUID",
		testProxyUsecaseFunc(&whitelistedCachedRequest, "project not found", false, projectNotFound, errors.New("no cache available"), nil, nil, projectNotFound))

	whitelistedNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns the proxy response When the path is not cacheable",
		testProxyUsecaseFunc(&whitelistedNotCachedRequest, "", true, nil, nil, nil, nil, nil))

	whitelistedCacheableNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, "127.0.0.1")
	t.Run("Returns no response When there is no cache and an proxy error",
		testProxyUsecaseFunc(&whitelistedCacheableNotCachedRequest, "no response from proxy", false, errors.New("no response from proxy"),
			errors.New("no cache available"), errors.New("proxy error"), nil, nil))
}

// Need docker-compose running in background
func TestProxyUsecase_Proxy_RedirectToMockServer_Integration(t *testing.T) {
	client := &http.Client{}
	newProject, _ := json.Marshal(inputs.NewProject{
		Title: "New Project",
	})

	// 0. Create a project
	req, err := http.NewRequest("POST", "http://0.0.0.0:8000/api/v1/projects", strings.NewReader(string(newProject)))
	if err != nil {
		t.Fatal(err)
	}

	r, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	_ = r.Body.Close()

	url, _ := r.Location()
	path := url.Path

	var uuidRegex = `(?m)([0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12})`
	re := regexp.MustCompile(uuidRegex)
	var uuid string
	for _, match := range re.FindAllString(path, -1) {
		uuid = match
	}

	// 1. Send a dummy request
	req, err = http.NewRequest("PUT", "http://0.0.0.0:8001/v1/"+uuid+"/mockserver/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	r, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = r.Body.Close()

	assert.Equal(t, string(b), `{
  "ports" : [ 1090 ],
  "version" : "5.9.0",
  "artifactId" : "mockserver-core",
  "groupId" : "org.mock-server"
}`)

	// 2. Get the number of request done by this project UUID
	req, err = http.NewRequest("GET", "http://0.0.0.0:8000/api/v1/projects/"+uuid, nil)
	if err != nil {
		t.Fatal(err)
	}

	r, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = r.Body.Close()

	type HTTPOutput struct {
		Data   outputs.ProjectOutputWithMetrics `json:"data"`
		Status string                           `json:"status"`
	}
	var projectOutputWithMetrics HTTPOutput
	err = json.Unmarshal(b, &projectOutputWithMetrics)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "New Project", projectOutputWithMetrics.Data.Title)
	assert.Equal(t, uuid, projectOutputWithMetrics.Data.UUID)
	assert.Equal(t, 1, projectOutputWithMetrics.Data.Metrics.RequestsCount)
	assert.Equal(t, 29, len(projectOutputWithMetrics.Data.Metrics.RequestsByDay))
	assert.Equal(t, 1, len(projectOutputWithMetrics.Data.Metrics.RPCUsage))
	assert.Equal(t, 1, projectOutputWithMetrics.Data.Metrics.RPCUsage[0].Value)
	assert.Equal(t, "/mockserver/status", projectOutputWithMetrics.Data.Metrics.RPCUsage[0].ID)
	assert.Equal(t, "/mockserver/status", projectOutputWithMetrics.Data.Metrics.RPCUsage[0].Label)
}

func testProxyUsecaseFunc(
	r *pkgmodel.Request,
	expResp string,
	expToRawProxy bool,
	expErr error,
	cacheErr error,
	proxyErr error,
	metricErr error,
	projectErr error) func(t *testing.T) {
	return func(t *testing.T) {
		// Given
		mockCache := stubBlockchainRepository([]byte("Dummy cache response"), cacheErr)
		mockProxy := stubBlockchainRepository([]byte("Dummy proxy response"), proxyErr)
		mockMetricsRepo := stubMetricsRepository(metricErr)
		mockCacheMetricsRepo := stubCacheMetricsRepository(nil)
		mockProjectRepo := stubProjectRepository(projectErr)
		puc := NewProxyUsecase(mockCache, mockProxy, mockMetricsRepo, mockProjectRepo, mockCacheMetricsRepo)

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

func stubMetricsRepository(error error) *mockMetricsRepository {
	mockMetricsRepository := &mockMetricsRepository{}
	mockMetricsRepository.
		On("Save", mock.Anything).
		Return(error).
		Once()
	mockMetricsRepository.
		On("SaveMany", mock.Anything).
		Return(error).
		Once()
	return mockMetricsRepository
}

func stubCacheMetricsRepository(error error) *mockCacheMetricsRepository {
	mockCacheMetricsRepository := &mockCacheMetricsRepository{}
	// TODO Add the stubs
	mockCacheMetricsRepository.
		On("Add", mock.Anything).
		Return(error).
		Once()
	mockCacheMetricsRepository.
		On("Len").
		Return(10).
		Once()

	noinputs := make([]*dbinputs.MetricsInput, 0)
	mockCacheMetricsRepository.
		On("GetAll").
		Return(noinputs, error).
		Once()
	return mockCacheMetricsRepository
}

func stubProjectRepository(error error) *mockProjectRepository {
	mockProjectRepository := &mockProjectRepository{}
	mockProjectRepository.
		On("FindByUUID", mock.Anything).
		Return(nil, error).
		Once()
	return mockProjectRepository
}
