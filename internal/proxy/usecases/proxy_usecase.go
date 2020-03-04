package usecases

import (
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	pkgrepository "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"time"
)

// ProxyUsecase contains the repositories and regexes to route paths proxying and store metrics
type ProxyUsecase struct {
	cacheRepo   repository.BlockchainRepository
	proxyRepo   repository.BlockchainRepository
	metricsRepo pkgrepository.MetricsRepository
	whitelisted []*regexp.Regexp
	blacklisted []*regexp.Regexp
	dontCache   []*regexp.Regexp
}

// ProxyUsecaseInterface contains all methods implemented by the proxyRepo use-case
type ProxyUsecaseInterface interface {
	Proxy(request *pkgmodel.Request) (response string, toRawProxy bool, err error)
}

// NoProxyResponse is the error message when there is no response from the proxyRepo
const NoProxyResponse = "no response from proxy"

// NewProxyUsecase returns a new proxy use-case
func NewProxyUsecase(
	cache repository.BlockchainRepository,
	proxy repository.BlockchainRepository,
	metricsRepo pkgrepository.MetricsRepository) *ProxyUsecase {
	return &ProxyUsecase{
		cacheRepo:   cache,
		proxyRepo:   proxy,
		metricsRepo: metricsRepo,
		whitelisted: setupRegexpFor(config.ProxyConfig.Proxy.WhitelistedMethods),
		blacklisted: setupRegexpFor(config.ProxyConfig.Proxy.BlockedMethods),
		dontCache:   setupRegexpFor(config.ProxyConfig.Proxy.DontCache),
	}
}

// Proxy proxy an http request to the right repositories
func (p *ProxyUsecase) Proxy(request *pkgmodel.Request) (response string, toRawProxy bool, err error) {
	logrus.Info("received proxy request for path: ", request.Path)
	r := []byte("call blacklisted")

	if !p.isAllowed(request.Path) {
		logrus.Debug("not allowed to proxy on the path: ", request.Path)
		return string(r), false, nil
	}

	if request.Action == pkgmodel.OBTAIN && p.isCacheable(request.Path) {
		r, err := p.cacheRepo.Get(request)
		if err != nil {
			logrus.Info("path not cached, fetching to node: ", request.Path)

			r, err = p.proxyRepo.Get(request)
			logrus.Info("received response from node: ", string(r.([]byte)))
			if err != nil {
				logrus.Errorf("could not request to proxy: %s", err)
				return errors.ErrNoProxyResponse.Error(), false, errors.ErrNoProxyResponse
			}

			_ = p.cacheRepo.Add(request, r)
		}

		// TODO first check if the project UUID is existing
		// TODO save the fact that it is cached from the LRU or not
		p.saveMetrics(request)
		return string(r.([]byte)), false, nil
	}

	p.saveMetrics(request)
	return "", true, nil
}

func (p *ProxyUsecase) saveMetrics(request *pkgmodel.Request) {
	metrics := inputs.NewMetricsInput(request, time.Now())
	err := p.metricsRepo.Save(&metrics)
	if err != nil {
		logrus.Errorf("could not save metrics: %s", err)
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
