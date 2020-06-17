package repository

import (
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
)

<<<<<<< HEAD
// TODO Rename
=======
>>>>>>> b6bf93d452ef683c38f71adbc215e1d7dafde8a6
// MetricInputRepository contains all available methods of a blockchain repository
type MetricInputRepository interface {
	Add(metric *inputs.MetricsInput) error
	GetAll() ([]*inputs.MetricsInput, error)
	Len() int
}
