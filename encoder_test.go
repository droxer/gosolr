package gosolr

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestJSONEncode(t *testing.T) {
	var doc = NewSolrDocument()
	doc.Add("name", "value", float32(1.0))

	var w bytes.Buffer
	encode([]*SolrDocument{doc}, &w)

	assert.Equal(t, `[{"name":"value"}]`, w.String())
}

func TestJSONDecode(t *testing.T) {
	var data = `
        {
          "responseHeader": {
            "status": 400,
            "QTime": 5
          },
          "error": {
            "msg": "Document is missing mandatory uniqueKey field: documentid",
            "code": 400
          }
        }
    `

	resp := &SolrResponse{}
	decode(strings.NewReader(data), resp)

	assert.Equal(t, 400, resp.Header.Status)
}
