package gosolr

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
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

		assert.Equal(t, `{"add":{"commitWithin":0,"overwrite":false,"boost":0,"doc":{"documentid":118523475}}}`, string(data))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(resp))
	}))
	defer ts.Close()

	solr := New(ts.URL+"/solr/collection1", time.Second*10)

	doc := &Document{
		Doc: map[string]interface{}{
			"documentid": 118523475,
		},
	}

	solrRep, _ := solr.Add(doc, true, false)
	assert.Equal(t, 200, solrRep.Header.Status)
}

func TestDeleteById(t *testing.T) {
	var resp = `
                {
                  "responseHeader": {
                    "status": 200,
                    "QTime": 0
                  }
                }
            `
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/solr/collection1/update?commit=true&wt=json", req.RequestURI)
		assert.Equal(t, "POST", req.Method)

		data, _ := ioutil.ReadAll(req.Body)

		assert.Equal(t, `{"delete":{"id":"123456"}}`, string(data))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(resp))
	}))
	defer ts.Close()

	solr := New(ts.URL+"/solr/collection1", time.Second*10)

	solrRep, _ := solr.DeleteById("123456", true)
	assert.Equal(t, 200, solrRep.Header.Status)
}

func TestCommit(t *testing.T) {
	var resp = `
                {
                  "responseHeader": {
                    "status": 200,
                    "QTime": 0
                  }
                }
            `
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/solr/collection1/update?commit=true&wt=json", req.RequestURI)
		assert.Equal(t, "GET", req.Method)

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(resp))
	}))
	defer ts.Close()

	solr := New(ts.URL+"/solr/collection1", time.Second*10)

	solrRep, _ := solr.Commit()
	assert.Equal(t, 200, solrRep.Header.Status)
}
