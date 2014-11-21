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
	var resp = `
                {
                  "responseHeader": {
                    "status": 200,
                    "QTime": 0
                  }
                }
            `
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/solr/collection1/update?commit=true&softCommit=false&wt=json", req.RequestURI)
		assert.Equal(t, "POST", req.Method)

		data, _ := ioutil.ReadAll(req.Body)

		assert.Equal(t, `[{"documentid":118523475}]`, string(data))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(resp))
	}))
	defer ts.Close()

	solr := NewSolr(ts.URL+"/solr/collection1", time.Second*10)
	doc := NewSolrDocument()
	doc.Add("documentid", 118523475, float32(1.0))

	solrRep, _ := solr.Add([]*SolrDocument{doc}, true, false)
	assert.Equal(t, 200, solrRep.Header.Status)
}
