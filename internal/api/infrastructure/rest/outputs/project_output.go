package outputs

import "github.com/octo-technology/tezos-link/backend/internal/api/domain/model"

// ProjectOutput contains the fields to represent a project
type ProjectOutput struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Key  string `json:"key"`
}

// NewProject returns a new project
func NewProjectOutputFromProject(project *model.Project) ProjectOutput {
    return ProjectOutput{
        ID:   project.ID,
        Name: project.Name,
        Key:  project.Key,
    }
}
