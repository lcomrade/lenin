[![Go report](https://goreportcard.com/badge/github.com/lcomrade/lenin)](https://goreportcard.com/report/github.com/lcomrade/lenin)
[![Go Reference](https://pkg.go.dev/badge/github.com/lcomrade/lenin.svg)](https://pkg.go.dev/github.com/lcomrade/lenin)
[![Release](https://img.shields.io/github/v/release/lcomrade/lenin)](https://github.com/lcomrade/lenin/releases/latest)
[![License](https://img.shields.io/github/license/lcomrade/lenin)](https://github.com/lcomrade/lenin/blob/main/LICENSE)

**Lenin** is a golang library for working with the Lenpaste API.

`NOTE:` Library releases correspond to the Lenpaste server releases.

## Install
Run: `go get github.com/lcomrade/lenin`

## Example
### Create new paste
```
package main

import(
	"github.com/lcomrade/lenin"
	"fmt"
)

func main() {
	// Base API URL
	baseURL := "https://paste.lcomrade.su/api"

	//Clean base URL
	baseURL = lenin.CleanURL(baseURL)

	// Create request
	req := lenin.NewReq {
		Title: "Lenin library",
		Text: "Lenin is a golang library for working with the Lenpaste API",
		Expiration: "5m",
	}

	// Do request
	resp, err := lenin.New(req, baseURL)
	if err != nil {
		panic(err)
	}

	// Print result
	fmt.Println("Paste name:", resp.Name)
}
```

## Documentation
- [Documentation](https://pkg.go.dev/github.com/lcomrade/lenin#section-documentation)
- [More examples](https://github.com/lcomrade/lenin/blob/main/docs/examples.md)
- [Popular errors](https://github.com/lcomrade/lenin/blob/main/docs/errors.md)
