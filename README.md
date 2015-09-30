# gosolr [![GoDoc](https://godoc.org/github.com/droxer/gosolr?status.svg)](https://godoc.org/github.com/droxer/gosolr)

[![Build Status](https://travis-ci.org/droxer/gosolr.svg?branch=master)](https://travis-ci.org/droxer/gosolr)

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
  solr := gosolr.New("http://localhost:8983/solr/collection1", time.Second*5)

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

## License

Copyright 2014

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.