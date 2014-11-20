package gosolr

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type JSONEncoder struct {
}

func (coder *JSONEncoder) Encode(docs []*SolrDocument, w io.Writer) (int, error) {
	var darray []map[string]interface{}

	for i := range docs {
		dmap := make(map[string]interface{})

		for k, v := range docs[i].Fields() {
			dmap[k] = v.Value()
		}

		darray = append(darray, dmap)
	}

	p, err := json.Marshal(darray)
	if err != nil {
		return -1, err
	}

	return w.Write(p)
}

func (coder *JSONEncoder) Decode(r io.Reader, to *SolrResponse) (int, error) {
	data, err := ioutil.ReadAll(r)
	json.Unmarshal(data, to)

	return len(data), err
}
