package gosolr

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

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

func (s *Solr) Add(docs []*SolrDocument, commit, softCommit bool) (*SolrResponse, error) {
	var (
		path    = "/update"
		headers = map[string]string{
			"Content-Type": "application/json",
		}
		params = map[string]string{
			"wt": "json",
		}
	)

	var buf bytes.Buffer
	encode(docs, &buf)
	params["commit"] = strconv.FormatBool(commit)
	params["softCommit"] = strconv.FormatBool(softCommit)

	return s.request("POST", path, headers, params, &buf)
}

func (s *Solr) DeleteById(id string, commit bool) (*SolrResponse, error) {
	return nil, nil
}

func (s *Solr) request(method, thePath string, headers, params map[string]string, buf *bytes.Buffer) (*SolrResponse, error) {

	req, err := http.NewRequest(method, s.url.String()+thePath+"?"+multimap(params).Encode(), buf)
	if err != nil {
		panic("error")
	}

	for k, v := range headers {
		req.Header[k] = []string{v}
	}

	resp, err := s.httpClient.Do(req)
	defer resp.Body.Close()
	if resp == nil {
		panic(err)
	}

	solrResp := &SolrResponse{}
	decode(resp.Body, solrResp)

	return solrResp, err
}
