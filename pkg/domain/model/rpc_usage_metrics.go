package model

// RPCUsageMetrics represents a number of requests for a given path
type RPCUsageMetrics struct {
	Path  string
	Value int
}

// NewRPCUsageMetrics returns a new RPCUsageMetrics
func NewRPCUsageMetrics(path string, value int) *RPCUsageMetrics {
	return &RPCUsageMetrics{
		Path:  path,
		Value: value,
	}
}
