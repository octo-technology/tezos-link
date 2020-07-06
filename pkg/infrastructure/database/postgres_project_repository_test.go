package database

import (
	"github.com/octo-technology/tezos-link/backend/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostgresProjectRepository_FindByUUID_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	pgr := NewPostgresProjectRepository(pg)
	creationDate := time.Now().Truncate(time.Millisecond).UTC()
	expectedProject := model.NewProject(1, "New Project", "A_KEY", creationDate, "CARTAGENET")
	_, err := pgr.Save(expectedProject.Title, expectedProject.UUID, creationDate, expectedProject.Network)
	if err != nil {
		t.Fatal(err)
	}

	// When
	p, err := pgr.FindByUUID("A_KEY")
	if err != nil {
		t.Fatal(err)
	}

	// Then
	assert.Equal(t, expectedProject.CreationDate.String(), p.CreationDate.String(), "Bad project")
	assert.Equal(t, expectedProject.UUID, p.UUID, "Bad project")
	assert.Equal(t, expectedProject.ID, p.ID, "Bad project")
	assert.Equal(t, expectedProject.Title, p.Title, "Bad project")
	assert.Equal(t, "Etc/UTC", p.CreationDate.Location().String(), "Bad project")
}

func TestPostgresProjectRepository_FindAll_Unit(t *testing.T) {
	// Given
	pool := getDockerPool()
	pg, resource := GetPostgresClient(*pool)
	defer pool.Purge(resource)

	creationDate := time.Now().Truncate(time.Millisecond).UTC()

	pgr := NewPostgresProjectRepository(pg)
	expectedFirstProject := model.NewProject(1, "New Project", "A_KEY", creationDate, "CARTAGENET")
	expectedSecondProject := model.NewProject(2, "New Project 2", "A_SECOND_KEY", creationDate.Add(2*time.Second).Truncate(time.Millisecond), "CARTAGENET")
	_, _ = pgr.Save(expectedFirstProject.Title, expectedFirstProject.UUID, creationDate, expectedFirstProject.Network)
	_, _ = pgr.Save(expectedSecondProject.Title, expectedSecondProject.UUID, creationDate.Add(2*time.Second).Truncate(time.Millisecond), expectedSecondProject.Network)

	// When
	p, err := pgr.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	// Then
	firstProject, secondProject := p[0], p[1]
	assert.Equal(t, expectedFirstProject.CreationDate.String(), firstProject.CreationDate.String(), "Bad project")
	assert.Equal(t, expectedFirstProject.UUID, firstProject.UUID, "Bad project")
	assert.Equal(t, expectedFirstProject.ID, firstProject.ID, "Bad project")
	assert.Equal(t, expectedFirstProject.Title, firstProject.Title, "Bad project")
	assert.Equal(t, expectedFirstProject.Network, firstProject.Network, "Bad project")
	assert.Equal(t, "Etc/UTC", firstProject.CreationDate.Location().String(), "Bad project")

	assert.Equal(t, expectedSecondProject.CreationDate.String(), secondProject.CreationDate.String(), "Bad project")
	assert.Equal(t, expectedSecondProject.UUID, secondProject.UUID, "Bad project")
	assert.Equal(t, expectedSecondProject.ID, secondProject.ID, "Bad project")
	assert.Equal(t, expectedSecondProject.Title, secondProject.Title, "Bad project")
	assert.Equal(t, expectedSecondProject.Network, secondProject.Network, "Bad project")
	assert.Equal(t, "Etc/UTC", secondProject.CreationDate.Location().String(), "Bad project")
}
