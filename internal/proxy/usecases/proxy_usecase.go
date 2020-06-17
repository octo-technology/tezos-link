package usecases

import (
	"regexp"
	"strings"
	"time"

	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	pkgrepository "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/sirupsen/logrus"
)

// ProxyUsecase contains the repositories and regexes to route paths proxying and store metrics
type ProxyUsecase struct {
	cacheRepo   repository.BlockchainRepository
	proxyRepo   repository.BlockchainRepository
	metricsRepo pkgrepository.MetricsRepository
	projectRepo pkgrepository.ProjectRepository
	lruMetrics  repository.MetricInputRepository
	whitelisted []*regexp.Regexp
	blacklisted []*regexp.Regexp
	dontCache   []*regexp.Regexp
}

// ProxyUsecaseInterface contains all methods implemented by the proxyRepo use-case
type ProxyUsecaseInterface interface {
	Proxy(request *pkgmodel.Request) (string, bool, error)
}

// NoProxyResponse is the error message when there is no response from the proxyRepo
const NoProxyResponse = "no response from proxy"

// NewProxyUsecase returns a new proxy use-case
func NewProxyUsecase(
	cacheRepo repository.BlockchainRepository,
	proxyRepo repository.BlockchainRepository,
	metricsRepo pkgrepository.MetricsRepository,
	projectRepo pkgrepository.ProjectRepository,
	lruMetrics repository.MetricInputRepository,
) *ProxyUsecase {

	return &ProxyUsecase{
		cacheRepo:   cacheRepo,
		proxyRepo:   proxyRepo,
		metricsRepo: metricsRepo,
		projectRepo: projectRepo,
		lruMetrics:  lruMetrics,
		whitelisted: setupRegexpFor(config.ProxyConfig.Proxy.WhitelistedMethods),
		blacklisted: setupRegexpFor(config.ProxyConfig.Proxy.BlockedMethods),
		dontCache:   setupRegexpFor(config.ProxyConfig.Proxy.DontCache),
	}
}

// TODO move this function to a separate usecase
func (p *ProxyUsecase) WriteCachedRequestsRoutine() {
	logrus.Info("func WriteCachedRequestsRoutine")
	allRequests, err := p.lruMetrics.GetAll()
	if err != nil {
		logrus.Error("could not init the LRU cache")
	}
	logrus.Info("len data", len(allRequests))
	err = p.metricsRepo.SaveMany(allRequests)
	if err != nil {
		logrus.Errorf("could not save metrics in database: %s", err)
	}

	time.Sleep(60 * time.Second)
}

// Proxy proxy an http request to the right repositories
func (p *ProxyUsecase) Proxy(request *pkgmodel.Request) (string, bool, error) {
	logrus.Info("received proxy request for path: ", request.Path)
	response := []byte("call blacklisted")

	_, err := p.projectRepo.FindByUUID(request.UUID)
	if err != nil {
		logrus.Debug("project ID not found: ", request.UUID, err.Error())
		return err.Error(), false, err
	}

	if !p.isAllowed(request.Path) {
		logrus.Debug("not allowed to proxy on the path: ", request.Path)
		return string(response), false, nil
	}

	if request.Action == pkgmodel.OBTAIN && p.isCacheable(request.Path) {
		response, err := p.cacheRepo.Get(request)
		if err != nil {
			logrus.Info("path not cached, fetching to node: ", request.Path)

			response, err = p.proxyRepo.Get(request)
			if err != nil {
				logrus.Errorf("could not request to proxy: %s", err)
				return errors.ErrNoProxyResponse.Error(), false, errors.ErrNoProxyResponse
			}
			logrus.Info("received response from node: ", string(response.([]byte)))

			_ = p.cacheRepo.Add(request, response)
		}

		// TODO save that it is cached from the LRU or not
		p.saveMetrics(request)
		return string(response.([]byte)), false, nil
	}

	p.saveMetrics(request)
	return "", true, nil
}

func (p *ProxyUsecase) saveMetrics(request *pkgmodel.Request) {
	metrics := inputs.NewMetricsInput(request, time.Now().UTC())

	// BEFORE
	//err := p.metricsRepo.Save(&metrics)

	// AFTER
	// add to cache
	err := p.lruMetrics.Add(&metrics) //TODO
	if err != nil {
		logrus.Error("could not init the LRU cache")
	}

	logrus.Info("metric input add to lruMetrics cache")
	// check limit reached if yes save in database
	nb := p.lruMetrics.Len()
	if nb >= config.ProxyConfig.Proxy.CacheMaxMetricItems {
		logrus.Info("metric top")
		allRequests, err := p.lruMetrics.GetAll()
		if err != nil {
			logrus.Error("could not init the LRU cache")
		}
		err = p.metricsRepo.SaveMany(allRequests) // TODO
		if err != nil {
			logrus.Errorf("could not save metrics in database: %s", err)
		}
	}

}

func (p *ProxyUsecase) isAllowed(url string) bool {
	ret := false
	urls := strings.Split(url, "?")
	url = "/" + strings.Trim(urls[0], "/")

	for _, wl := range p.whitelisted {
		if wl.Match([]byte(url)) {
			ret = true
			for _, bl := range p.blacklisted {
				if bl.Match([]byte(url)) {
					ret = false
					break
				}
			}
			break
		}
	}

	return ret
}

func (p *ProxyUsecase) isCacheable(url string) bool {
	ret := true

	for _, wl := range p.dontCache {
		if wl.Match([]byte(url)) {
			ret = false
		}
	}

	return ret
}

func setupRegexpFor(regexPaths []string) []*regexp.Regexp {
	var list []*regexp.Regexp

	for _, s := range regexPaths {
		regex, err := regexp.Compile(s)
		if err != nil {
			logrus.Error("could not compile Regexp: ", s)
		} else {
			list = append(list, regex)
		}
	}

	return list
}
