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
    baseUrl string
    client *http.Client
}

func NewProxyBlockchainRepository() repository.BlockchainRepository {
    baseUrl := "http://" + config.ProxyConfig.Tezos.Host + ":" + strconv.Itoa(config.ProxyConfig.Tezos.Port)
    client := &http.Client{Timeout: time.Duration(config.ProxyConfig.Proxy.ReadTimeout) * time.Second}

    return &proxyBlockchainRepository{
        baseUrl: baseUrl,
        client:  client,
    }
}

func (p proxyBlockchainRepository) Get(request *model.Request) (interface{}, error) {
    url := p.baseUrl + request.Path

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        logrus.Error(fmt.Sprintf("Error while proxying to blockchain node: %s", err))
        return nil, err
    }

    r, err := p.client.Do(req)
    var b []byte
    b, err = ioutil.ReadAll(r.Body)
    if err != nil {
        logrus.Error(fmt.Sprintf("Error getting response from blockchain node: %s", err))
        return nil, err
    }
    _ = r.Body.Close()

    return b, nil
}

func (p proxyBlockchainRepository) Add(request *model.Request, response interface{}) error {
    panic("not implemented")
}

