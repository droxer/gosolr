# gosolr

Apache Solr client implementation by go.

This is **draft** implementation for now. Contribute is more than welcome.

## Example Usage

```go
package main

import (
  "fmt"
  "github.com/droxer/gosolr"
  "log"
  "time"
)

func main() {
  solr := gosolr.NewSolr("http://localhost:8983/solr/collection1", time.Second*5)

  var doc = &gosolr.Document{
    Doc: map[string]interface{}{
      "documentid":  118523475,
      "audiences":   []string{"gosolr"},
      "collapse_id": 8888,
    },
  }

  resp, _ := solr.Add(doc, true, false)

  log.Print(fmt.Sprintf("Status: %v", resp.Header.Status))
}


```