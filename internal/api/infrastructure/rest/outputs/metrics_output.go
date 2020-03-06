package outputs

import (
	"fmt"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
)

// MetricsOutput contains the fields to represent a metrics
type MetricsOutput struct {
	RequestsCount int              `json:"requestsCount"`
	RequestsByDay []*RequestsByDay `json:"requestsByDay"`
}

// NewMetricsOutput returns a new metrics
func NewMetricsOutput(metrics *model.Metrics) MetricsOutput {
	var requestsByDay []*RequestsByDay
	for i := 0; i < len(metrics.RequestsByDay); i++ {
		requestsByDay = append(requestsByDay, NewRequestsByDay(metrics.RequestsByDay[i]))
	}

	return MetricsOutput{
		RequestsCount: metrics.RequestsCount,
		RequestsByDay: requestsByDay,
	}
}

// RequestsByDay contains the date and value field to represents requests for a given day
type RequestsByDay struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

// NewRequestsByDay returns a new RequestsByDay
func NewRequestsByDay(requestsByDay *pkgmodel.RequestsByDayMetrics) *RequestsByDay {
	return &RequestsByDay{
		Date:  fmt.Sprintf("%s-%s-%s", requestsByDay.Year, requestsByDay.Month, requestsByDay.Day),
		Value: requestsByDay.Value,
	}
}
