package model

// Project contains the fields to represent a project
type Project struct {
	ID   int64
	Name string
	Key  string
}

// NewProject returns a new project
func NewProject(ID int64, name string, key string) Project {
	return Project{
		ID:   ID,
		Name: name,
		Key:  key,
	}
}
