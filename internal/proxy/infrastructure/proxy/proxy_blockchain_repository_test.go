package proxy

import (
	"github.com/bmizerany/assert"
	"github.com/octo-technology/tezos-link/backend/config"
	"testing"
)

func TestProxyblockchainRepository_IsRollingRedirection_Unit(t *testing.T) {
	// Given
	_, err := config.ParseProxyConf("../../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}

	proxyBC := NewProxyBlockchainRepository()
	requestedURL := "/chains/main/blocks/head"

	// When
	response_bool := proxyBC.IsRollingRedirection(requestedURL)
	if err != nil {
		t.Fatal("could not save in cache", err)
	}

	// Then
	assert.Equal(t, response_bool, true)
}

func TestProxyblockchainRepository_IsRollingRedirection_Archive_Unit(t *testing.T) {
	// Given
	_, err := config.ParseProxyConf("../../../../test/proxy/conf/test.toml")
	if err != nil {
		t.Fatal("could not parse conf", err)
	}

	proxyBC := NewProxyBlockchainRepository()
	requestedURL := "/chains/main/blocks/min_date"

	// When
	response_bool := proxyBC.IsRollingRedirection(requestedURL)
	if err != nil {
		t.Fatal("could not save in cache", err)
	}

	// Then
	assert.Equal(t, response_bool, false)
}
