package display

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertHtmlToText(t *testing.T) {
	golds := []struct {
		input  string
		output string
	}{
		{"<html>test1</html>", "test1"},
		{"test2", "test2"},
	}

	assert.Equal(t, golds[0].output, ConvertHtmlToText(golds[0].input))
	assert.Equal(t, golds[1].output, ConvertHtmlToText(golds[1].input))
}
