package repository

import "github.com/octo-technology/tezos-link/backend/pkg/domain/model"

// MetricRepository contains all available methods of the metric repository
type MetricRepository interface {
	Save(metric *model.Metric) error
	Count(uuid string) (int, error)
}
