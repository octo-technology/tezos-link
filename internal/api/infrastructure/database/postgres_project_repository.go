package database

import (
	"database/sql"
	"fmt"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	"github.com/sirupsen/logrus"
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
	rows, err := pg.connection.Query("SELECT id, title, uuid FROM projects")
	if err != nil {
		return nil, fmt.Errorf("no projects found: %s", err)
	}

	var r []*model.Project
	for rows.Next() {
		cur := model.Project{}
		err := rows.Scan(&cur.ID, &cur.Title, &cur.UUID)
		if err != nil {
			return nil, fmt.Errorf("could not map projects: %s", err)
		}
		r = append(r, &cur)
	}

	return r, nil
}

// FindByUUID finds a project by uuid
func (pg *postgresProjectRepository) FindByUUID(uuid string) (*model.Project, error) {
	r := model.Project{}
	err := pg.connection.
		QueryRow("SELECT id, title, uuid FROM projects WHERE uuid = $1", uuid).
		Scan(&r.ID, &r.Title, &r.UUID)

	if err != nil {
		logrus.Errorf("project %s not found: %s", uuid, err)
		return nil, errors.ErrProjectNotFound
	}

	return &r, nil
}

// Save insert a new project
func (pg postgresProjectRepository) Save(title string, uuid string) (*model.Project, error) {
	r := model.Project{}

	err := pg.connection.
		QueryRow("INSERT INTO projects(title, uuid) VALUES ($1, $2) RETURNING id, title, uuid", title, uuid).
		Scan(&r.ID, &r.Title, &r.UUID)

	if err != nil {
		return nil, fmt.Errorf("could not insert project %s: %s", title, err)
	}

	return &r, nil
}

// Ping ping the database
func (pg *postgresProjectRepository) Ping() error {
	err := pg.connection.Ping()
	if err != nil {
		return err
	}

	return nil
}
