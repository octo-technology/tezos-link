package database

import (
    "database/sql"
    "errors"
    "fmt"
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/model"
    "github.com/octo-technology/tezos-link/backend/internal/backend/domain/repository"
)

type postgresProjectRepository struct {
    connection *sql.DB
}

func NewPostgresProjectRepository(connection *sql.DB) repository.ProjectRepository {
    return &postgresProjectRepository{
        connection: connection,
    }
}

func (pg postgresProjectRepository) FindAll() ([]*model.Project, error) {
    rows, err := pg.connection.Query("SELECT id, name, key FROM projects")
    if err != nil {
        return nil, errors.New(fmt.Sprintf("no projects found: %s", err))
    }

    var r []*model.Project
    for rows.Next() {
        cur := model.Project{}
        err := rows.Scan(&cur.ID, &cur.Name, &cur.Key)
        if err != nil {
            return nil, errors.New(fmt.Sprintf("could not map projects: %s", err))
        }
        r = append(r, &cur)
    }

    return r, nil
}

func (pg postgresProjectRepository) FindById(id int64) (*model.Project, error) {
    r := model.Project{}
    err := pg.connection.
        QueryRow("SELECT id, name, key FROM projects WHERE id = $1", id).
        Scan(&r.ID, &r.Name, &r.Key)

    if err != nil {
        return nil, errors.New(fmt.Sprintf("project %d not found: %s", id, err))
    }

    return &r, nil
}

func (pg postgresProjectRepository) Save(name string, key string) (*model.Project, error) {
    r := model.Project{}

    err := pg.connection.
        QueryRow("INSERT INTO projects(name, key) VALUES ($1, $2) RETURNING id, name, key", name, key).
        Scan(&r.ID, &r.Name, &r.Key)

    if err != nil {
        return nil, errors.New(fmt.Sprintf("could not insert project %s: %s", name, err))
    }

    return &r, nil
}

func (pg postgresProjectRepository) UpdateById(project *model.Project) error {
    panic("implement me")
}

func (pg postgresProjectRepository) DeleteById(project *model.Project) error {
    panic("implement me")
}

func (pg postgresProjectRepository) Ping() error {
    err := pg.connection.Ping()
    if err != nil {
        return err
    }

    return nil
}
