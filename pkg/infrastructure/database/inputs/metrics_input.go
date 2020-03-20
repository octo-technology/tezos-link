package inputs

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"time"
)

// MetricsInput represents a new metric insert for a request made on a given date
type MetricsInput struct {
	Request *model.Request
	Date    time.Time
}

// NewMetricsInput returns a new MetricsInput
func NewMetricsInput(request *model.Request, date time.Time) MetricsInput {
	return MetricsInput{
		Request: request,
		Date:    date,
	}
}
