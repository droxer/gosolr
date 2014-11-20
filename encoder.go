package gosolr

import (
	"io"
)

type Encoder interface {
	Encode(docs []*SolrDocument, w io.Writer) (int, error)
	Decode(r io.Reader, to *SolrResponse) (int, error)
}
