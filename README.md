## Installation
```
go get github.com/AaronHarden/jsontodb
```
## Code Example
```
package main

import (
	"database/sql"

  "github.com/AaronHarden/jsontodb"
	_ "github.com/lib/pq"
)

func main() {
  file, err := os.Open("dbconn.json")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  db, err = jsontodb.JSONToPQ(file)
  //
}
```

## Unofficially Supported Drivers
pq - https://github.com/lib/pq

## Tests
```
go test
```
