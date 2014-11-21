package gosolr

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddDoc(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/solr/collection1/update?commit=true&softCommit=false&wt=json", req.RequestURI)
		assert.Equal(t, "POST", req.Method)

		data, _ := ioutil.ReadAll(req.Body)

		assert.Equal(t, `[{"documentid":118523475}]`, string(data))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("ok"))
	}))
	defer ts.Close()

	solr := NewSolr(ts.URL+"/solr/collection1", time.Second*10)
	doc := NewSolrDocument()
	doc.Add("documentid", 118523475, float32(1.0))

	solr.Add([]*SolrDocument{doc}, true, false)
}

func TestSolrRep(t *testing.T) {
	var resp = `
                {
                  "responseHeader": {
                    "status": 400,
                    "QTime": 1
                  },
                  "error": {
                    "msg": "Document is missing mandatory uniqueKey field: documentid",
                    "code": 400
                  }
                }
            `

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(resp))
	}))
	defer ts.Close()

	server := NewSolr(ts.URL+"/solr/collection1", time.Second*10)

	doc := NewSolrDocument()
	doc.Add("documentid", 118523475, float32(1.0))

	solrRep, _ := server.Add([]*SolrDocument{doc}, true, true)

	assert.Equal(t, "Document is missing mandatory uniqueKey field: documentid", solrRep.Error.Msg)
}
