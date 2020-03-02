package repository

import (
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
)

// MetricsRepository contains all available methods of the metrics repository
type MetricsRepository interface {
	Save(metrics *inputs.MetricsInput) error
	Count(uuid string) (int, error)
}
