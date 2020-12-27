> This repository is in **Work in Progress** state. If you need something, create an [issue](https://github.com/ondrejsika/poste-go/issues/new)

# poste-go - Go client for Poste.io API

    2020 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/poste-go

## Example Usage

```go
package main

import (
	"fmt"

	"github.com/ondrejsika/poste-go"
)

func main() {
	api := poste.New("http://localhost/admin/api", "admin@example.com", "asdfasdf")
	api.CreateDomain("foo.io")
	api.CreateBox("foo@foo.io", "asdfasdf")
  sieve, _ := api.GetBoxSieve("foo@foo.io")
  fmt.Println(sieve)
}
```
