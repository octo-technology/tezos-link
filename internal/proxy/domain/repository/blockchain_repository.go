package repository

import "github.com/octo-technology/tezos-link/backend/internal/proxy/domain/model"

// BlockchainRepository contains all available methods of a blockchain repository
type BlockchainRepository interface {
	Get(request *model.Request) (interface{}, error)
	Add(request *model.Request, response interface{}) error
}
