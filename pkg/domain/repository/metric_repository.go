package repository

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/octo-technology/tezos-link/backend/pkg/infrastructure/database/inputs"
	"time"
)

// MetricsRepository contains all available methods of the metrics repository
type MetricsRepository interface {
	Save(metrics *inputs.MetricsInput) error
	CountAll(uuid string) (int, error)
	FindRequestsByDay(uuid string, from time.Time, to time.Time) ([]*model.RequestsByDayMetrics, error)
}
