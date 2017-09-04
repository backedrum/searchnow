package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchHN(t *testing.T) {
	terms := []string{"newstories", "topstories", "beststories", "askstories", "jobstories"}
	for _, term := range terms {
		results := fetchHackerNews(term, 5)
		assert.True(t, len(results) > 0)
		for _, result := range results {
			assert.NotEmpty(t, result.Title)
		}
	}
}
