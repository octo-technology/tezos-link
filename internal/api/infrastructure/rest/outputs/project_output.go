package outputs

import "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"

// ProjectOutputWithMetrics contains the fields to represent a project
type ProjectOutputWithMetrics struct {
	Name    string        `json:"name"`
	UUID    string        `json:"uuid"`
	Metrics MetricsOutput `json:"metrics"`
}

// NewProjectOutputWithMetrics returns a new project with metrics
func NewProjectOutputWithMetrics(project *model.Project, metrics *model.Metrics) ProjectOutputWithMetrics {
	return ProjectOutputWithMetrics{
		Name:    project.Name,
		UUID:    project.UUID,
		Metrics: NewMetricsOutput(metrics),
	}
}
