package gosolr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiMap(t *testing.T) {
	var params = map[string]string{
		"lang":    "en",
		"country": "us",
	}

	assert.Equal(t, "country=us&lang=en", multimap(params).Encode())
}
