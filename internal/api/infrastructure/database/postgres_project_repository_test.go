package database

import (
	"github.com/octo-technology/tezos-link/backend/internal/api/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresProjectRepository_FindByUUID_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresProjectRepository(pg)
	expectedProject := model.NewProject(1, "New Project", "A_KEY")
	s, err := pgr.Save(expectedProject.Name, expectedProject.UUID)
	if err != nil {
		t.Fatal(err)
	}

	// When
	p, err := pgr.FindByUUID("A_KEY")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, &expectedProject, p, "Bad project")
	assert.Equal(t, &expectedProject, s, "Bad project")
}

func TestPostgresProjectRepository_FindAll_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresProjectRepository(pg)
	expectedFirstProject := model.NewProject(1, "New Project", "A_KEY")
	expectedSecondProject := model.NewProject(2, "New Project 2", "A_SECOND_KEY")
	_, _ = pgr.Save(expectedFirstProject.Name, expectedFirstProject.UUID)
	_, _ = pgr.Save(expectedSecondProject.Name, expectedSecondProject.UUID)

	// When
	p, err := pgr.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	// Then
	firstProject, secondProject := p[0], p[1]
	assert.Equal(t, &expectedFirstProject, firstProject, "Bad project")
	assert.Equal(t, &expectedSecondProject, secondProject, "Bad project")
}
