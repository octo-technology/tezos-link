package proxy

import (
	"fmt"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type proxyBlockchainRepository struct {
	client *http.Client
}

// NewProxyBlockchainRepository returns a new blockchain proxy repository
func NewProxyBlockchainRepository() repository.BlockchainRepository {

	client := &http.Client{Timeout: time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second}

	return &proxyBlockchainRepository{
		client: client,
	}
}

func (p proxyBlockchainRepository) Get(request *pkgmodel.Request, url string) (interface{}, error) {
	// redirect to Archive nodes by default

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
