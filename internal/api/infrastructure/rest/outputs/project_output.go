package outputs

import "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"

// ProjectOutputWithMetrics contains the fields to represent a project
type ProjectOutputWithMetrics struct {
	Title   string        `json:"title"`
	UUID    string        `json:"uuid"`
	Metrics MetricsOutput `json:"metrics"`
}

// NewProjectOutputWithMetrics returns a new project with metrics
func NewProjectOutputWithMetrics(project *model.Project, metrics *model.Metrics) ProjectOutputWithMetrics {
	return ProjectOutputWithMetrics{
		Title:   project.Title,
		UUID:    project.UUID,
		Metrics: NewMetricsOutput(metrics),
	}
}
