package gosolr

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Document struct {
	CommitWithin int32                  `json:"commitWithin"`
	Overwrite    bool                   `json:"overwrite"`
	Boost        float32                `json:"boost"`
	Doc          map[string]interface{} `json:"doc"`
}

type Solr struct {
	url        *url.URL
	httpClient http.Client
}

func NewSolr(connectionString string, timeout time.Duration) *Solr {
	url, err := url.Parse(connectionString)
	if err != nil {
		panic(err)
	}

	server := &Solr{
		url: url,
	}

	transport := &http.Transport{DisableKeepAlives: false}
	transport.ResponseHeaderTimeout = timeout
	server.httpClient.Transport = transport
	return server
}

func (s *Solr) Search(query string, params map[string]string) (*SolrResponse, error) {
	var (
		path    = "/select"
		buf     bytes.Buffer
		headers = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
		}
	)

	params["wt"] = "json"
	params["q"] = query
	buf.WriteString("")
	return s.request("GET", path, headers, params, &buf)
}

func (s *Solr) Add(doc *Document, commit, softCommit bool) (*SolrResponse, error) {
	var (
		path    = "/update"
		buf     bytes.Buffer
		headers = map[string]string{
			"Content-Type": "application/json",
		}
		params = map[string]string{
			"wt": "json",
		}
		addReq = map[string]*Document{
			"add": doc,
		}
	)

	b, err := json.Marshal(addReq)
	if err != nil {
		log.Panicln(err.Error())
	}

	buf.Write(b)
	params["commit"] = strconv.FormatBool(commit)
	params["softCommit"] = strconv.FormatBool(softCommit)

	return s.request("POST", path, headers, params, &buf)
}

func (s *Solr) DeleteById(id string, commit bool) (*SolrResponse, error) {
	var (
		path    = "/update"
		buf     bytes.Buffer
		headers = map[string]string{
			"Content-Type": "application/json",
		}
		params = map[string]string{
			"wt": "json",
		}
		delReq = make(map[string]interface{})
	)

	delReq["delete"] = map[string]string{
		"id": id,
	}

	data, _ := json.Marshal(delReq)
	buf.WriteString(string(data))

	params["commit"] = strconv.FormatBool(commit)
	return s.request("POST", path, headers, params, &buf)
}

func (s *Solr) Commit() (*SolrResponse, error) {
	var (
		path    = "/update"
		buf     bytes.Buffer
		headers = map[string]string{
			"Content-Type": "application/json",
		}
		params = map[string]string{
			"wt": "json",
		}
	)

	params["commit"] = "true"
	return s.request("GET", path, headers, params, &buf)
}

func (s *Solr) request(method, thePath string, headers, params map[string]string, buf *bytes.Buffer) (*SolrResponse, error) {
	var requestUrl = s.url.String() + thePath + "?" + multimap(params).Encode()
	log.Println(requestUrl)

	req, err := http.NewRequest(method, requestUrl, buf)
	if err != nil {
		log.Panicln(err.Error())
	}

	for k, v := range headers {
		req.Header[k] = []string{v}
	}

	resp, err := s.httpClient.Do(req)
	defer resp.Body.Close()
	if resp == nil {
		log.Panicln(err.Error())
	}

	solrResp := &SolrResponse{}
	data, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(data, solrResp)
	return solrResp, err
}
