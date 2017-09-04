package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchIpLocation(t *testing.T) {
	results := searchIpLocation("8.8.8.8", -1)

	assert.Equal(t, 1, len(results))
	assert.Equal(t, 1, len(results[0].ExtrasOrder))

	extra := results[0].ExtrasOrder[0]
	assert.NotEmpty(t, extra)
	assert.Equal(t, 1, len(results[0].Extras))
	assert.NotEmpty(t, results[0].Extras[extra])
}

func TestIsValidIp4(t *testing.T) {
	golds := []struct {
		ipstring string
		isValid  bool
	}{
		{"9.56.24.100", true},
		{"2001:0db8:0a0b:12f0:0000:0000:0000:0001", false},
		{"not an ip, obviously", false},
	}

	for _, gold := range golds {
		assert.Equal(t, gold.isValid, isValidIp4(gold.ipstring))
	}
}
