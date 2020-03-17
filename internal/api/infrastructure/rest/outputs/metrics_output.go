package outputs

import (
	"fmt"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
)

// MetricsOutput contains the fields to represent a metrics
type MetricsOutput struct {
	RequestsCount int              `json:"requestsCount"`
	RequestsByDay []*RequestsByDay `json:"requestsByDay"`
	RPCUsage      []*RPCUsage      `json:"rpcUsage"`
}

// NewMetricsOutput returns a new metrics
func NewMetricsOutput(metrics *pkgmodel.Metrics) MetricsOutput {
	requestsByDay := make([]*RequestsByDay, len(metrics.RequestsByDay))
	for i := 0; i < len(metrics.RequestsByDay); i++ {
		requestsByDay[i] = NewRequestsByDay(metrics.RequestsByDay[i])
	}

	rpcUsage := make([]*RPCUsage, len(metrics.RPCUSage))
	for i := 0; i < len(metrics.RPCUSage); i++ {
		rpcUsage[i] = NewRPCUsage(metrics.RPCUSage[i])
	}

	return MetricsOutput{
		RequestsCount: metrics.RequestsCount,
		RequestsByDay: requestsByDay,
		RPCUsage:      rpcUsage,
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

// RPCUsage contains the path and value to have some metrics on rpc usage
type RPCUsage struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value int    `json:"value"`
}

// NewRPCUsage returns a new RPCUsage
func NewRPCUsage(rpcUsage *pkgmodel.RPCUsageMetrics) *RPCUsage {
	return &RPCUsage{
		ID:    rpcUsage.Path,
		Label: rpcUsage.Path,
		Value: rpcUsage.Value,
	}
}
