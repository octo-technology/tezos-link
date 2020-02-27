package proxy

import (
	"fmt"
	"github.com/octo-technology/tezos-link/backend/config"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/proxy/domain/repository"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type proxyBlockchainRepository struct {
	baseURL string
	client  *http.Client
}

// NewProxyBlockchainRepository returns a new blockchain proxy repository
func NewProxyBlockchainRepository() repository.BlockchainRepository {
	baseURL := "http://" + config.ProxyConfig.Tezos.Host + ":" + strconv.Itoa(config.ProxyConfig.Tezos.Port)
	client := &http.Client{Timeout: time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second}

	return &proxyBlockchainRepository{
		baseURL: baseURL,
		client:  client,
	}
}

func (p proxyBlockchainRepository) Get(request *model.Request) (interface{}, error) {
	url := p.baseURL + request.Path

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

func (p proxyBlockchainRepository) Add(request *model.Request, response interface{}) error {
	panic("not implemented")
}
