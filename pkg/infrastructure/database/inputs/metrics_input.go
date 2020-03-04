package inputs

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"time"
)

type MetricsInput struct {
	Request *model.Request
	Date    time.Time
}

func NewMetricsInput(request *model.Request, date time.Time) MetricsInput {
	return MetricsInput{
		Request: request,
		Date:    date,
	}
}
