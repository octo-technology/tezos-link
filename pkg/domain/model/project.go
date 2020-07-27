package model

import "time"

// Project contains the fields to represent a project
type Project struct {
	ID           int64
	Title        string
	UUID         string
	CreationDate time.Time
	Network      string
}

// NewProject returns a new project
func NewProject(ID int64, title string, uuid string, creationDate time.Time, network string) Project {
	return Project{
		ID:           ID,
		Title:        title,
		UUID:         uuid,
		CreationDate: creationDate,
		Network:      network,
	}
}
