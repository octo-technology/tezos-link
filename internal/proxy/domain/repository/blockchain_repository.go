package repository

import (
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
)

// BlockchainRepository contains all available methods of a blockchain repository
type BlockchainRepository interface {
	Get(request *pkgmodel.Request) (interface{}, error)
	Add(request *pkgmodel.Request, response interface{}) error
}
