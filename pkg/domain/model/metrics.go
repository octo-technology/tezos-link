package model

// Metrics contains the fields to represent project's metrics
type Metrics struct {
	RequestsCount int
	RequestsByDay []*RequestsByDayMetrics
	RPCUSage      []*RPCUsageMetrics
}

// NewMetrics returns a new metrics
func NewMetrics(requestsCount int, requestsByDay []*RequestsByDayMetrics, rpcUsage []*RPCUsageMetrics) Metrics {
	return Metrics{
		RequestsCount: requestsCount,
		RequestsByDay: requestsByDay,
		RPCUSage:      rpcUsage,
	}
}
