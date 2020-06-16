package cache

import (
	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	"testing"
	"time"
)

func TestLRUProjectRepository_FindByUUID_Unit(t *testing.T) {
	// Given
	_, err := config.ParseProxyConf("../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}

	projectCache := NewLRUProjectRepository()
	projectUUID := "12234"
	projectTitle := "DummyProject"
	projectCreationDate := time.Now()

	// When
	_, err = projectCache.Save(projectTitle, projectUUID, projectCreationDate)
	if err != nil {
		t.Fatal("could not save in cache", err)
	}

	prj, err := projectCache.FindByUUID(projectUUID)
	if err != nil {
		t.Fatal("project not found in cache", err)
	}

	// Then
	assert.Equal(t, prj.UUID, projectUUID)
	assert.Equal(t, prj.Title, projectTitle)
	assert.Equal(t, prj.CreationDate, projectCreationDate)

}
