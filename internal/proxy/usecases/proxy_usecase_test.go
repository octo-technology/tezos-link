package usecases

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/inputs"
	"github.com/octo-technology/tezos-link/backend/internal/api/infrastructure/rest/outputs"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	dbinputs "github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/stretchr/testify/mock"
)

type dummyResponse struct {
	expResp                 string
	expToRawProxy           bool
	expErr                  error
	cacheErr                error
	proxyErr                error
	metricErr               error
	projectErr              error
	projectCacheErr         error
	projectDatabaseResponse *pkgmodel.Project
}

func newDummyResponse() dummyResponse {
	return dummyResponse{
		expResp:                 "",
		expToRawProxy:           false,
		expErr:                  nil,
		cacheErr:                nil,
		proxyErr:                nil,
		metricErr:               nil,
		projectErr:              nil,
		projectCacheErr:         nil,
		projectDatabaseResponse: nil,
	}
}

func TestProxyUsecase_Proxy_Unit(t *testing.T) {
	_, err := config.ParseProxyConf("../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}
	prj := pkgmodel.NewProject(123, "DUMMY_TITLE", "DUMMY_UUID", time.Now(), "CARTHAGENET")
	localURL := "127.0.0.1"

	blockedRequest := pkgmodel.NewRequest("/dummy/path", "UUID", pkgmodel.OBTAIN, localURL)
	expResponse := newDummyResponse()
	expResponse.expResp = "call blacklisted"
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns blacklisted When there is a blacklisted path",
		testProxyUsecaseFunc(&blockedRequest, &expResponse))

	postRequest := pkgmodel.NewRequest("/chains/main/blocks/head", "UUID", pkgmodel.PUSH, localURL)
	expResponse = newDummyResponse()
	expResponse.expToRawProxy = true
	expResponse.projectDatabaseResponse = &prj
	t.Run("Forward to reverse proxy When there is a PUSH request",
		testProxyUsecaseFunc(&postRequest, &expResponse))

	whitelistedCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, localURL)
	expResponse = newDummyResponse()
	expResponse.expResp = "Dummy cache response"
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns the cached response When there is a whitelisted path",
		testProxyUsecaseFunc(&whitelistedCachedRequest, &expResponse))

	expResponse = newDummyResponse()
	expResponse.expResp = "Dummy proxy response"
	expResponse.cacheErr = errors.New("no cache available")
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns the proxy response When there is no cached response",
		testProxyUsecaseFunc(&whitelistedCachedRequest, &expResponse))

	projectNotFound := errors.New("project not found")
	expResponse = newDummyResponse()
	expResponse.expResp = "project not found"
	expResponse.expErr = projectNotFound
	expResponse.cacheErr = errors.New("no cache available")
	expResponse.projectErr = projectNotFound
	expResponse.projectCacheErr = projectNotFound
	t.Run("Returns no project found error When it is not in the cache or database",
		testProxyUsecaseFunc(&whitelistedCachedRequest, &expResponse))

	expResponse = newDummyResponse()
	expResponse.expResp = "Dummy proxy response"
	expResponse.cacheErr = errors.New("no cache available")
	expResponse.projectCacheErr = projectNotFound
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns project When it is not in the cache but found in database",
		testProxyUsecaseFunc(&whitelistedCachedRequest, &expResponse))

	whitelistedNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks", "UUID", pkgmodel.OBTAIN, localURL)
	expResponse = newDummyResponse()
	expResponse.expToRawProxy = true
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns the proxy response When the path is not cacheable",
		testProxyUsecaseFunc(&whitelistedNotCachedRequest, &expResponse))

	whitelistedCacheableNotCachedRequest := pkgmodel.NewRequest("/chains/main/blocks/number", "UUID", pkgmodel.OBTAIN, localURL)
	expResponse = newDummyResponse()
	expResponse.expResp = "no response from proxy"
	expResponse.expErr = errors.New("no response from proxy")
	expResponse.cacheErr = errors.New("no cache available")
	expResponse.proxyErr = errors.New("proxy error")
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns no response When there is no cache and an proxy error",
		testProxyUsecaseFunc(&whitelistedCacheableNotCachedRequest, &expResponse))

	redirectionRequest := pkgmodel.NewRequest("/chains/main/blocks/head", "UUID", pkgmodel.OBTAIN, localURL)
	expResponse = newDummyResponse()
	expResponse.expResp = "Dummy proxy response(redirection)"
	expResponse.cacheErr = errors.New("no cache available")
	expResponse.projectDatabaseResponse = &prj
	t.Run("Returns no response When there is no cache and an proxy error",
		testProxyUsecaseFunc(&redirectionRequest, &expResponse))
}

func testProxyUsecaseFunc(req *pkgmodel.Request, resp *dummyResponse) func(t *testing.T) {
	return func(t *testing.T) {
		rollingUrlPattern := "http://" + config.ProxyConfig.Tezos.RollingHost + ":" + strconv.Itoa(config.ProxyConfig.Tezos.RollingPort)
		// Given
		mockCache := stubBlockchainRepository([]byte("Dummy cache response"), resp.cacheErr)
		mockProxy := stubProxyBlockchainRepository([]byte("Dummy proxy response"), resp.proxyErr, rollingUrlPattern, []byte("Dummy proxy response(redirection)"))
		mockMetricsRepo := stubMetricsRepository(resp.metricErr)
		mockProjectRepo := stubProjectRepository(resp.projectDatabaseResponse, resp.projectErr)
		mockProjectCacheRepo := stubProjectRepository(nil, resp.projectCacheErr)
		mockCacheMetricsRepo := stubCacheMetricsRepository(nil)

		puc := NewProxyUsecase(mockCache, mockProxy, mockMetricsRepo, mockProjectRepo, mockProjectCacheRepo, mockCacheMetricsRepo)

		// When
		proxyResp, toRawProxy, err := puc.Proxy(req)

		// Then
		assert.Equal(t, resp.expResp, proxyResp)
		assert.Equal(t, resp.expToRawProxy, toRawProxy)
		assert.Equal(t, resp.expErr, err)
	}
}

func stubProxyBlockchainRepository(response interface{}, error error, pattern string, responseRedirection interface{}) *mockBlockchainRepository {
	mockBlockchainRepository := &mockBlockchainRepository{}
	mockBlockchainRepository.
		On("Get", mock.Anything, mock.MatchedBy(func(url string) bool { return strings.Contains(url, pattern) })).
		Return(responseRedirection, error).
		Once()
	mockBlockchainRepository.
		On("Get", mock.Anything, mock.Anything).
		Return(response, error).
		Once()
	mockBlockchainRepository.
		On("Add", mock.Anything, mock.Anything).
		Return(error).
		Once()
	return mockBlockchainRepository
}

func stubBlockchainRepository(response interface{}, error error) *mockBlockchainRepository {
	mockBlockchainRepository := &mockBlockchainRepository{}
	mockBlockchainRepository.
		On("Get", mock.Anything, mock.Anything).
		Return(response, error).
		Once()
	mockBlockchainRepository.
		On("Add", mock.Anything, mock.Anything).
		Return(error).
		Once()
	return mockBlockchainRepository
}

func stubProjectRepository(response *pkgmodel.Project, error error) *mockProjectRepository {
	mockProjectRepository := &mockProjectRepository{}
	mockProjectRepository.
		On("FindByUUID", mock.Anything).
		Return(response, error).
		Once()
	mockProjectRepository.
		On("Save", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, error).
		Once()
	return mockProjectRepository
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

	now := time.Now().UTC()
	nowMinusOneMonth := now.AddDate(0, -1, 0)
	rangeTime := now.Sub(nowMinusOneMonth)
	nbDaysThisMonth := int(rangeTime.Hours() / 24)

	assert.Equal(t, "New Project", projectOutputWithMetrics.Data.Title)
	assert.Equal(t, uuid, projectOutputWithMetrics.Data.UUID)
	assert.Equal(t, 1, projectOutputWithMetrics.Data.Metrics.RequestsCount)
	assert.Equal(t, nbDaysThisMonth, len(projectOutputWithMetrics.Data.Metrics.RequestsByDay))
	assert.Equal(t, 1, len(projectOutputWithMetrics.Data.Metrics.RPCUsage))
	assert.Equal(t, 1, projectOutputWithMetrics.Data.Metrics.RPCUsage[0].Value)
	assert.Equal(t, "/mockserver/status", projectOutputWithMetrics.Data.Metrics.RPCUsage[0].ID)
	assert.Equal(t, "/mockserver/status", projectOutputWithMetrics.Data.Metrics.RPCUsage[0].Label)
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
