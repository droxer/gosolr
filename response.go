package gosolr

type SolrResponse struct {
	Header   *Header   `json:"responseHeader"`
	Response *Response `json:"error,omitempty"`
	Error    Error     `json:"error,omitempty"`
}

type Header struct {
	Status int
	QTime  int
	Params *Params `json:"params,omitempty"`
}

type Response struct {
}

type Error struct {
	Msg  string
	Code int
}

type Params struct {
	Indent bool
	Q      string
	_      string
	wt     string
}
