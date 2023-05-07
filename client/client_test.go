package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client/transport"
)

func ExampleNewClient() {
	var err error
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return
	}

	api := suggest.Api{
		Client: NewClient(endpointUrl),
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

func ExampleCredentials() {
	var err error
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return
	}

	creds := Credentials{
		ApiKeyValue:    "<YOUR_API_KEY>",
		SecretKeyValue: "<YOUR_SECRET_KEY>",
	}

	api := suggest.Api{
		Client: NewClient(endpointUrl, WithCredentialProvider(&creds)),
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

func ExampleWithEncoderFactory() {
	var err error
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return
	}

	// Customize json encoding
	encoderFactory := func(w io.Writer) transport.Encoder {
		e := json.NewEncoder(w)
		e.SetIndent("", "  ")
		return func(v interface{}) error {
			return e.Encode(v)
		}
	}

	// Customize json decoding
	decoderFactory := func(r io.Reader) transport.Decoder {
		d := json.NewDecoder(r)
		d.DisallowUnknownFields()
		return func(v interface{}) error {
			return d.Decode(v)
		}
	}

	api := suggest.Api{
		Client: NewClient(endpointUrl, WithEncoderFactory(encoderFactory), WithDecoderFactory(decoderFactory)),
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
