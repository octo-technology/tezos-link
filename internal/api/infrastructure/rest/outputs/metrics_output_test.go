package outputs

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMetricsOutput_Unit(t *testing.T) {
	// Given
	firstMetrics := pkgmodel.NewRequestsByDayMetrics("2020", "03", "05", 4)
	secondMetrics := pkgmodel.NewRequestsByDayMetrics("2019", "03", "06", 5)
	rpcUsage := pkgmodel.NewRPCUsageMetrics("/dummy/path", 3)
	metrics := model.NewMetrics(
		3,
		[]*pkgmodel.RequestsByDayMetrics{firstMetrics, secondMetrics},
		[]*pkgmodel.RPCUsageMetrics{rpcUsage})

	// When
	metricsOutput := NewMetricsOutput(&metrics)

	// Then
	assert.Equal(t, 3, metricsOutput.RequestsCount)
	assert.Equal(t, 4, metricsOutput.RequestsByDay[0].Value)
	assert.Equal(t, "2020-03-05", metricsOutput.RequestsByDay[0].Date)
	assert.Equal(t, 5, metricsOutput.RequestsByDay[1].Value)
	assert.Equal(t, "2019-03-06", metricsOutput.RequestsByDay[1].Date)
	assert.Equal(t, "/dummy/path", metricsOutput.RPCUsage[0].Label)
	assert.Equal(t, "/dummy/path", metricsOutput.RPCUsage[0].ID)
	assert.Equal(t, 3, metricsOutput.RPCUsage[0].Value)
}
