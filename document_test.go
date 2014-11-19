package gosolr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddDuplicatedFiled(t *testing.T) {
	doc := NewSolrDocument()
	doc.Add("name", "foo", float32(1.0))
	doc.Add("name", "bar", float32(1.0))

	assert.Equal(t, "bar", doc.Get("name").Value().(string))
}
