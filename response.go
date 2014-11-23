package gosolr

type Term []interface{}

type SolrResponse struct {
	Header Header          `json:"responseHeader,omitempty"`
	Terms  map[string]Term `json:"terms,omitempty"`
	Result Result          `json:"response,omitempty"`
	Error  Error           `json:"error,omitempty"`
}

type Header struct {
	Status int
	QTime  int
	Params Params `json:"params,omitempty"`
}

type Doc map[string]interface{}

type Result struct {
	NumFound int   `json:"numFound"`
	Start    int   `json:"start"`
	Docs     []Doc `json:"docs"`
}

type Error struct {
	Msg  string
	Code int
}

type Params struct {
	Indent bool
	Q      string
	WT     string
}
