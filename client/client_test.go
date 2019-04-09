package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ekomobile/dadata/v2/api/suggest"
)

func ExampleNewClient() {
	var err error
	cleanEndpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return
	}

	api := suggest.Api{
		Client: NewClient(cleanEndpointUrl),
	}

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

func ExampleNewClient_Credentials() {
	var err error
	cleanEndpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return
	}

	creds := Credentials{
		ApiKeyValue:    "<YOUR_API_KEY>",
		SecretKeyValue: "<YOUR_API_KEY>",
	}

	api := suggest.Api{
		Client: NewClient(cleanEndpointUrl, WithCredentialProvider(&creds)),
	}

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
