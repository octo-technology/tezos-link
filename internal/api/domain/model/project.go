package model

// Project contains the fields to represent a project
type Project struct {
	ID    int64
	Title string
	UUID  string
}

// NewProject returns a new project
func NewProject(ID int64, title string, uuid string) Project {
	return Project{
		ID:    ID,
		Title: title,
		UUID:  uuid,
	}
}
