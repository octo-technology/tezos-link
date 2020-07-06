package usecases

import (
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestProjectUsecase_FindProjectAndMetrics_Unit(t *testing.T) {
	// Given
	p := pkgmodel.NewProject(123, "A PROJECT", "AN_UUID", time.Now().UTC(), "CARTAGENET")
	str := "2014-11-12T11:45:26.371Z"
	now, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Fatal(err)
	}

	firstRPCUsage := pkgmodel.NewRPCUsageMetrics("/dummy/path", 3)
	secondRPCUsage := pkgmodel.NewRPCUsageMetrics("/dummy/another", 10)
	stubRPCUSageMetrics := []*pkgmodel.RPCUsageMetrics{firstRPCUsage, secondRPCUsage}
	firstRequestMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "1", 4)
	secondRequestMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "12", 5)
	stubRequestsMetrics := []*pkgmodel.RequestsByDayMetrics{firstRequestMetrics, secondRequestMetrics}
	expRequestsMetrics := []*pkgmodel.RequestsByDayMetrics{
		firstRequestMetrics,
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "2", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "3", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "4", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "5", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "6", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "7", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "8", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "9", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "10", 0),
		pkgmodel.NewRequestsByDayMetrics("2014", "11", "11", 0),
		secondRequestMetrics,
	}
	stubLastRequests := []string{
		"/chains/main/blocks/head/header",
		"/chains/main/blocks/head/header/head",
	}

	expMetrics := pkgmodel.NewMetrics(2, expRequestsMetrics, stubRPCUSageMetrics, stubLastRequests)

	mockProjectRepository := &mockProjectRepository{}
	mockProjectRepository.
		On("FindByUUID", mock.Anything).
		Return(&p, nil).
		Once()

	mockMetricsRepository := &mockMetricsRepository{}
	mockMetricsRepository.
		On("CountAll", mock.Anything).
		Return(2, nil).
		Once()

	mockMetricsRepository.
		On("FindRequestsByDay", mock.Anything, mock.Anything, mock.Anything).
		Return(stubRequestsMetrics, nil).
		Once()

	mockMetricsRepository.
		On("CountRPCPathUsage", mock.Anything, mock.Anything, mock.Anything).
		Return(stubRPCUSageMetrics, nil).
		Once()

	mockMetricsRepository.
		On("FindLastRequests", mock.Anything).
		Return(stubLastRequests, nil).
		Once()

	pu := NewProjectUsecase(mockProjectRepository, mockMetricsRepository)

	// When
	projects, metrics, err := pu.FindProjectAndMetrics("AN_UUID", now.AddDate(0, 0, -12), now)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, &p, projects)
	assert.ElementsMatch(t, expMetrics.RequestsByDay, metrics.RequestsByDay)
	assert.ElementsMatch(t, expMetrics.RPCUSage, metrics.RPCUSage)
	assert.ElementsMatch(t, stubLastRequests, metrics.LastRequests)
}
