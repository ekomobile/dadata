# Client for DaData.ru

Forked from https://github.com/webdeskltd/dadata.


[![Build Status](https://travis-ci.org/ekomobile/dadata.svg)](https://travis-ci.org/ekomobile/dadata)
[![GitHub release](https://img.shields.io/github/release/ekomobile/dadata.svg)](https://github.com/ekomobile/dadata/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ekomobile/dadata)](https://goreportcard.com/report/github.com/ekomobile/dadata)
[![GoDoc](https://godoc.org/github.com/ekomobile/dadata?status.svg)](https://godoc.org/github.com/ekomobile/dadata)

DaData API v2

Implemented [Clean](https://dadata.ru/api/clean/) and [Suggest](https://dadata.ru/api/suggest/) methods.

## Installation

`go get github.com/ekomobile/dadata/v2`

## Usage
```go
import (
	"context"
	"fmt"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/suggest"
)

func DaDataExample()  {
	api := dadata.NewSuggestApi()

	params := suggest.RequestParams{
		Query: "ул Свободы",
	}

	suggestions, err := api.Address(context.Background(), &params)
	if err != nil {
		return
	}

	for _, s := range suggestions {
		fmt.Printf("%s", s.Value)
	}
}
```

## Licence
MIT see [LICENSE](LICENSE)
