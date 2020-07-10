package database

import (
	"database/sql"
	"fmt"
	"github.com/octo-technology/tezos-link/backend/pkg/domain/errors"
	pkgmodel "github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	pkgrepository "github.com/octo-technology/tezos-link/backend/pkg/domain/repository"
	"github.com/sirupsen/logrus"
	"time"
)

type postgresProjectRepository struct {
	connection *sql.DB
}

// NewPostgresProjectRepository returns a new postgres project repository
func NewPostgresProjectRepository(connection *sql.DB) pkgrepository.ProjectRepository {
	return &postgresProjectRepository{
		connection: connection,
	}
}

// FindAll returns all projects
func (pg postgresProjectRepository) FindAll() ([]*pkgmodel.Project, error) {
	rows, err := pg.connection.Query("SELECT id, title, uuid, creation_date, network FROM projects")
	if err != nil {
		return nil, fmt.Errorf("no projects found: %s", err)
	}

	var projects []*pkgmodel.Project
	for rows.Next() {
		cursor := pkgmodel.Project{}
		err := rows.Scan(&cursor.ID, &cursor.Title, &cursor.UUID, &cursor.CreationDate, &cursor.Network)
		if err != nil {
			return nil, fmt.Errorf("could not map projects: %s", err)
		}
		projects = append(projects, &cursor)
	}

	return projects, nil
}

// FindByUUID finds a project by uuid
func (pg *postgresProjectRepository) FindByUUID(uuid string) (*pkgmodel.Project, error) {
	stmt, err := pg.connection.Prepare("SELECT id, title, uuid, creation_date, network FROM projects WHERE uuid = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	project := pkgmodel.Project{}
	err = stmt.
		QueryRow(uuid).
		Scan(&project.ID, &project.Title, &project.UUID, &project.CreationDate, &project.Network)

	if err != nil {
		logrus.Errorf("project %s not found: %s", uuid, err)
		return nil, errors.ErrProjectNotFound
	}

	return &project, nil
}

// Save insert a new project
func (pg postgresProjectRepository) Save(title string, uuid string, creationDate time.Time, network string) (*pkgmodel.Project, error) {
	project := pkgmodel.Project{}

	err := pg.connection.
		QueryRow("INSERT INTO projects(title, uuid, creation_date, network) VALUES ($1, $2, $3, $4) RETURNING id, title, uuid, creation_date, network", title, uuid, creationDate, network).
		Scan(&project.ID, &project.Title, &project.UUID, &project.CreationDate, &project.Network)

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
