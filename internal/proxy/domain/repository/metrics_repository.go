package repository

import (
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
)

// TODO Rename
// MetricInputRepository contains all available methods of a blockchain repository
type MetricInputRepository interface {
	Add(metric *inputs.MetricsInput) error
	GetAll() ([]*inputs.MetricsInput, error)
	Len() int
}
