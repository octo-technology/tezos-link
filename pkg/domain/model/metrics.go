package model

// Metrics contains the fields to represent project's metrics
type Metrics struct {
	RequestsCount int
	RequestsByDay []*RequestsByDayMetrics
	RPCUSage      []*RPCUsageMetrics
	Last10RPC     []string
}

// NewMetrics returns a new metrics
func NewMetrics(
	requestsCount int,
	requestsByDay []*RequestsByDayMetrics,
	rpcUsage []*RPCUsageMetrics,
	last10RPC []string) Metrics {
	return Metrics{
		RequestsCount: requestsCount,
		RequestsByDay: requestsByDay,
		RPCUSage:      rpcUsage,
		Last10RPC:     last10RPC,
	}
}
