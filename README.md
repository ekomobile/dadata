# Client for DaData.ru

Forked from https://github.com/webdeskltd/dadata.


[![Build Status](https://travis-ci.org/ekomobile/dadata.svg)](https://travis-ci.com/ekomobile/dadata)
[![GitHub release](https://img.shields.io/github/release/ekomobile/dadata.svg)](https://github.com/ekomobile/dadata/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ekomobile/dadata)](https://goreportcard.com/report/github.com/ekomobile/dadata/v2)
[![GoDoc](https://godoc.org/github.com/ekomobile/dadata/v2?status.svg)](https://godoc.org/github.com/ekomobile/dadata/v2)

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


## Configuration 

### Credentials

`DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables are used by default to authenticate client.

Custom credential provider may be used by implementing `client.CredentialProvider` interface.

Also, there is a "simple" credential provider `client.Credentials` you may utilize.

```go
creds := client.Credentials{
    ApiKeyValue:    "<YOUR_API_KEY>",
    SecretKeyValue: "<YOUR_SECRET_KEY>",
}

api := NewSuggestApi(client.WithCredentialProvider(&creds))
```


### HTTP client

HTTP client may be overridden with custom one:

```go
httpClient := &http.Client{}

api := NewSuggestApi(WithHttpClient(httpClient))
```


## Licence
MIT see [LICENSE](LICENSE)
