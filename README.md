# gosolr

Apache Solr client implementation by go.

This is **draft** implementation for now. Contribute is more than welcome.

## Example Usage

```go
package main

import (
  "github.com/droxer/gosolr"
  "log"
  "fmt"
)

func main() {
   solr := NewSolr("http://localhost:8983/solr/collection1", time.Second*5)
   doc := NewSolrDocument()
   doc.Add("documentid", "7777", float32(1.0))
   doc.Add("audiences", []string{"droxer"}, float32(1.0))
   doc.Add("collapse_id", "8888", float32(1.0))
   resp, _ := solr.Add([]*SolrDocument{doc}, true, false)

   log.Print(fmt.Sprintf("Status: %v", resp.Header.Status))
}


```