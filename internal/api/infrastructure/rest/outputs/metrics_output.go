package outputs

import "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"

// MetricsOutput contains the fields to represent a metrics
type MetricsOutput struct {
	RequestsCount int `json:"requestsCount"`
}

// NewMetricsOutput returns a new metrics
func NewMetricsOutput(metrics *model.Metrics) MetricsOutput {
	return MetricsOutput{
		RequestsCount: metrics.RequestsCount,
	}
}
