package usecases

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestProjectUsecase_FindProjectAndMetrics_Unit(t *testing.T) {
	// Given
	p := model.NewProject(123, "A PROJECT", "AN_UUID")
	str := "2014-11-12T11:45:26.371Z"
	now, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Fatal(err)
	}

	firstMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "1", 4)
	secondMetrics := pkgmodel.NewRequestsByDayMetrics("2014", "11", "12", 5)
	stubMetrics := []*pkgmodel.RequestsByDayMetrics{firstMetrics, secondMetrics}
	expRequestsMetrics := []*pkgmodel.RequestsByDayMetrics{
		firstMetrics,
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
		secondMetrics,
	}
	expMetrics := model.NewMetrics(2, expRequestsMetrics)

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
		Return(stubMetrics, nil).
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
}
