package database

import (
	"database/sql"
	"fmt"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type postgresProjectRepository struct {
	connection *sql.DB
}

// NewPostgresProjectRepository returns a new postgres project repository
func NewPostgresProjectRepository(connection *sql.DB) repository.ProjectRepository {
	return &postgresProjectRepository{
		connection: connection,
	}
}

// FindAll returns all projects
func (pg postgresProjectRepository) FindAll() ([]*model.Project, error) {
	rows, err := pg.connection.Query("SELECT id, title, uuid, creation_date FROM projects")
	if err != nil {
		return nil, fmt.Errorf("no projects found: %s", err)
	}

	var projects []*model.Project
	for rows.Next() {
		cursor := model.Project{}
		err := rows.Scan(&cursor.ID, &cursor.Title, &cursor.UUID, &cursor.CreationDate)
		if err != nil {
			return nil, fmt.Errorf("could not map projects: %s", err)
		}
		projects = append(projects, &cursor)
	}

	return projects, nil
}

// FindByUUID finds a project by uuid
func (pg *postgresProjectRepository) FindByUUID(uuid string) (*model.Project, error) {
	project := model.Project{}
	err := pg.connection.
		QueryRow("SELECT id, title, uuid, creation_date FROM projects WHERE uuid = $1", uuid).
		Scan(&project.ID, &project.Title, &project.UUID, &project.CreationDate)

	if err != nil {
		logrus.Errorf("project %s not found: %s", uuid, err)
		return nil, errors.ErrProjectNotFound
	}

	return &project, nil
}

// Save insert a new project
func (pg postgresProjectRepository) Save(title string, uuid string, creationDate time.Time) (*model.Project, error) {
	project := model.Project{}

	err := pg.connection.
		QueryRow("INSERT INTO projects(title, uuid, creation_date) VALUES ($1, $2, $3) RETURNING id, title, uuid, creation_date", title, uuid, creationDate).
		Scan(&project.ID, &project.Title, &project.UUID, &project.CreationDate)

	if err != nil {
		return nil, fmt.Errorf("could not insert project %s: %s", title, err)
	}

	return &project, nil
}

// Ping ping the database
func (pg *postgresProjectRepository) Ping() error {
	err := pg.connection.Ping()
	if err != nil {
		return err
	}

	return nil
}
