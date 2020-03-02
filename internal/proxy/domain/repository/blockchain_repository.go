package repository

import (
	model2 "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
)

// BlockchainRepository contains all available methods of a blockchain repository
type BlockchainRepository interface {
	Get(request *model2.Request) (interface{}, error)
	Add(request *model2.Request, response interface{}) error
}
