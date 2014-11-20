package gosolr

type SolrResponse struct {
	Header struct {
		Status int
		QTime  int
	} `json:"responseHeader"`

	Error struct {
		Msg  string
		Code int
	} `json:"error"`
}
