package usecases

import (
	"fmt"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

// ProxyUsecase contains the repositories and regexes to route paths proxying
type ProxyUsecase struct {
	cache       repository.BlockchainRepository
	proxy       repository.BlockchainRepository
	whitelisted []*regexp.Regexp
	blacklisted []*regexp.Regexp
	dontCache   []*regexp.Regexp
}

// ProxyUsecaseInterface contains all methods implemented by the proxy use-case
type ProxyUsecaseInterface interface {
	Proxy(request *model.Request) (response string, toRawProxy bool, err error)
}

// NoProxyResponse is the error message when there is no response from the proxy
const NoProxyResponse = "no response from proxy"

// NewProxyUsecase returns a new proxy use-case
func NewProxyUsecase(cache repository.BlockchainRepository, proxy repository.BlockchainRepository) *ProxyUsecase {
	return &ProxyUsecase{
		cache:       cache,
		proxy:       proxy,
		whitelisted: setupRegexpFor(config.ProxyConfig.Proxy.WhitelistedMethods),
		blacklisted: setupRegexpFor(config.ProxyConfig.Proxy.BlockedMethods),
		dontCache:   setupRegexpFor(config.ProxyConfig.Proxy.DontCache),
	}
}

// Proxy proxy an http request to the right repositories
func (p *ProxyUsecase) Proxy(request *model.Request) (response string, toRawProxy bool, err error) {
	logrus.Info("received proxy request for path: ", request.Path)
	r := []byte("call blacklisted")

	if !p.isAllowed(request.Path) {
		logrus.Debug("not allowed to proxy on the path: ", request.Path)
		return string(r), false, nil
	}

	if request.Action == model.OBTAIN && p.isCacheable(request.Path) {
		r, err := p.cache.Get(request)
		if err != nil {
			logrus.Info("path not cached, fetching to node: ", request.Path)

			r, err = p.proxy.Get(request)
			logrus.Info("received response from node: ", string(r.([]byte)))
			if err != nil {
				return NoProxyResponse, false, fmt.Errorf("could not request to proxy: %s", err)
			}

			_ = p.cache.Add(request, r)
		}

		// TODO: Insert to metrics
		return string(r.([]byte)), false, nil
	}

	// TODO: Insert to metrics
	return "", true, nil
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
