package outputs

// MetricOutput contains the fields to represent a metric
type MetricOutput struct {
	UUID          string `json:"uuid"`
	RequestsCount int    `json:"requestsCount"`
}
