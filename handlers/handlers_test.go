package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasEngineSupport(t *testing.T) {
	golds := []struct {
		engineName string
		supported  bool
	}{
		{"google", true},
		{"so", true},
		{"geocities", false},
	}

	for _, gold := range golds {
		result := HasEngineSupport(gold.engineName)
		assert.Equal(t, result, gold.supported)
	}
}
