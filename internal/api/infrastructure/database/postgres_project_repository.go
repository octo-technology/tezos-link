package database

import (
	"database/sql"
	"fmt"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/repository"
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
	rows, err := pg.connection.Query("SELECT id, name, key FROM projects")
	if err != nil {
		return nil, fmt.Errorf("no projects found: %s", err)
	}

	var r []*model.Project
	for rows.Next() {
		cur := model.Project{}
		err := rows.Scan(&cur.ID, &cur.Name, &cur.Key)
		if err != nil {
			return nil, fmt.Errorf("could not map projects: %s", err)
		}
		r = append(r, &cur)
	}

	return r, nil
}

// FindByID finds a project by id
func (pg postgresProjectRepository) FindByID(id int64) (*model.Project, error) {
	r := model.Project{}
	err := pg.connection.
		QueryRow("SELECT id, name, key FROM projects WHERE id = $1", id).
		Scan(&r.ID, &r.Name, &r.Key)

	if err != nil {
		return nil, fmt.Errorf("project %d not found: %s", id, err)
	}

	return &r, nil
}

// Save save a new project
func (pg postgresProjectRepository) Save(name string, key string) (*model.Project, error) {
	r := model.Project{}

	err := pg.connection.
		QueryRow("INSERT INTO projects(name, key) VALUES ($1, $2) RETURNING id, name, key", name, key).
		Scan(&r.ID, &r.Name, &r.Key)

	if err != nil {
		return nil, fmt.Errorf("could not insert project %s: %s", name, err)
	}

	return &r, nil
}

// UpdateByID update a project by id
func (pg postgresProjectRepository) UpdateByID(project *model.Project) error {
	panic("implement me")
}

// DeleteByID delete a project by id
func (pg postgresProjectRepository) DeleteByID(project *model.Project) error {
	panic("implement me")
}

// Ping ping the database
func (pg postgresProjectRepository) Ping() error {
	err := pg.connection.Ping()
	if err != nil {
		return err
	}

	return nil
}
