package model

import "github.com/octo-technology/tezos-link/backend/pkg/domain/model"

// Metrics contains the fields to represent project's metrics
type Metrics struct {
	RequestsCount int
	RequestsByDay []*model.RequestsByDayMetrics
	RPCUSage      []*model.RPCUsageMetrics
}

// NewMetrics returns a new metrics
func NewMetrics(requestsCount int, requestsByDay []*model.RequestsByDayMetrics, rpcUsage []*model.RPCUsageMetrics) Metrics {
	return Metrics{
		RequestsCount: requestsCount,
		RequestsByDay: requestsByDay,
		RPCUSage:      rpcUsage,
	}
}
