# Client for DaData.ru

Forked from https://github.com/webdeskltd/dadata.


[![Build Status](https://travis-ci.org/ekomobile/dadata.svg)](https://travis-ci.org/ekomobile/dadata)
[![GitHub release](https://img.shields.io/github/release/ekomobile/dadata.svg)](https://github.com/ekomobile/dadata/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ekomobile/dadata)](https://goreportcard.com/report/github.com/ekomobile/dadata)
[![GoDoc](https://godoc.org/github.com/ekomobile/dadata?status.svg)](https://godoc.org/github.com/ekomobile/dadata)

DaData API v2

Implemented [Clean](https://dadata.ru/api/clean/) and [Suggest](https://dadata.ru/api/suggest/) methods.

## Installation

`go get github.com/ekomobile/dadata`

## Usage
```go
package main

import (
	"fmt"

	"github.com/ekomobile/dadata"
)

func main() {
	// By default client gets keys from `DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables.
	daData := dadata.NewClient()

	banks, err := daData.SuggestBanks(dadata.SuggestRequestParams{Query: "Кредитный", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, bank := range banks {
		fmt.Println(bank.Data.Name.Full)
		fmt.Println(bank.Data.Bic)
	}

	// Output:
	// "МОСКОВСКИЙ КРЕДИТНЫЙ БАНК" (ПУБЛИЧНОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО)
	// 044525659
	// КОММЕРЧЕСКИЙ БАНК "РЕСПУБЛИКАНСКИЙ КРЕДИТНЫЙ АЛЬЯНС" (ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ)
	// 044525860
	// ЖИЛИЩНО-КРЕДИТНЫЙ КОММЕРЧЕСКИЙ БАНК "ЖИЛКРЕДИТ" ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ
	// 044525325
}
```

## Licence
MIT see [LICENSE](LICENSE)
