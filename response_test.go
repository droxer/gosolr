package gosolr

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rawResp = `
    {
        "responseHeader":{
            "status":0,
            "QTime":0,
            "params":{
                "q":"documentid:22222222",
                "wt":"json"
                }
        },
        "response":{
            "numFound":1,
            "start":0,
            "docs":[
                {
                    "audiences":["gosolr"],
                    "collapse_id":"8888",
                    "documentid":"22222222"
                }
            ]
        }
    }
`

func TestResponseStruct(t *testing.T) {
	var resp = &SolrResponse{}
	json.Unmarshal([]byte(rawResp), resp)

	assert.Equal(t, 0, resp.Header.Status)
	assert.Equal(t, 0, resp.Header.QTime)
	assert.Equal(t, "documentid:22222222", resp.Header.Params.Q)
	assert.Equal(t, "json", resp.Header.Params.WT)

	assert.Equal(t, 1, resp.Result.NumFound)
	assert.Equal(t, 0, resp.Result.Start)
	assert.Equal(t, 1, len(resp.Result.Docs))
	assert.Equal(t, "22222222", resp.Result.Docs[0]["documentid"])
	assert.Equal(t, "8888", resp.Result.Docs[0]["collapse_id"])
	assert.Equal(t, []interface{}{"gosolr"}, resp.Result.Docs[0]["audiences"])
}
