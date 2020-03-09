package model

type RPCUsageMetrics struct {
	Path  string
	Value int
}

func NewRPCUsageMetrics(path string, value int) *RPCUsageMetrics {
	return &RPCUsageMetrics{
		Path:  path,
		Value: value,
	}
}
