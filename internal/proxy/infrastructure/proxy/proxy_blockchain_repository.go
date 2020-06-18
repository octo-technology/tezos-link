package proxy

import (
	"fmt"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type proxyBlockchainRepository struct {
	baseArchiveURL 	string
	baseRollingURL 	string
	client         	*http.Client
	rollingPatterns []*regexp.Regexp
}

// NewProxyBlockchainRepository returns a new blockchain proxy repository
func NewProxyBlockchainRepository() repository.BlockchainRepository {
	baseArchiveURL := "http://" + config.ProxyConfig.Tezos.ArchiveHost + ":" + strconv.Itoa(config.ProxyConfig.Tezos.Port)
	baseRollingURL := "http://" + config.ProxyConfig.Tezos.RollingHost + ":" + strconv.Itoa(config.ProxyConfig.Tezos.RollingPort)
	client := &http.Client{Timeout: time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second}

	return &proxyBlockchainRepository{
		baseArchiveURL: baseArchiveURL,
		baseRollingURL: baseRollingURL,
		client:         client,
		rollingPatterns: setupRegexpFor(config.ProxyConfig.Tezos.WhitelistedRolling),
	}
}

func (p proxyBlockchainRepository) Get(request *pkgmodel.Request) (interface{}, error) {
	// redirect to Archive nodes by default
	url := p.baseArchiveURL + request.Path
	//if strings.Contains(request.Path, "/block/head") {
	if p.IsRollingRedirection(request.Path) {
		url = p.baseRollingURL + request.Path
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error while building request for blockchain node: %s", err))
		return nil, err
	}

	r, err := p.client.Do(req)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error while requesting to blockchain node: %s", err))
		return nil, err
	}

	var b []byte
	b, err = ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error while reading response from blockchain node: %s", err))
		return nil, err
	}
	_ = r.Body.Close()

	logrus.Info(url, b)
	return b, nil
}

func (p proxyBlockchainRepository) Add(request *pkgmodel.Request, response interface{}) error {
	panic("not implemented")
}

func (p proxyBlockchainRepository) IsRollingRedirection(url string) bool {
	ret := false
	urls := strings.Split(url, "?")
	url = "/" + strings.Trim(urls[0], "/")

	for _, wl := range p.rollingPatterns {
		if wl.Match([]byte(url)) {
			ret = true
			break
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