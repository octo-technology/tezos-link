package outputs

import (
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
)

// ProjectOutputWithMetrics contains the fields to represent a project
type ProjectOutputWithMetrics struct {
	Title   string        `json:"title"`
	UUID    string        `json:"uuid"`
	Network string        `json:"network"`
	Metrics MetricsOutput `json:"metrics"`
}

// NewProjectOutputWithMetrics returns a new project with metrics
func NewProjectOutputWithMetrics(project *pkgmodel.Project, metrics *pkgmodel.Metrics) ProjectOutputWithMetrics {
	return ProjectOutputWithMetrics{
		Title:   project.Title,
		UUID:    project.UUID,
		Network: project.Network,
		Metrics: NewMetricsOutput(metrics),
	}
}
