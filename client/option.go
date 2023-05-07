package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ekomobile/dadata/v2/client/transport"
)

type (
	// Option applies some option to clientOptions.
	Option func(opts *clientOptions)
)

// WithHttpClient sets custom http.Client.
func WithHttpClient(c *http.Client) Option {
	return func(opts *clientOptions) {
		opts.httpClient = c
	}
}

// WithCredentialProvider sets credential provider.
func WithCredentialProvider(c CredentialProvider) Option {
	return func(opts *clientOptions) {
		opts.credentialProvider = c
	}
}

func WithEncoderFactory(f transport.EncoderFactory) Option {
	return func(opts *clientOptions) {
		opts.encoderFactory = f
	}
}

func WithDecoderFactory(f transport.DecoderFactory) Option {
	return func(opts *clientOptions) {
		opts.decoderFactory = f
	}
}

func applyOptions(options *clientOptions, opts ...Option) {
	for _, o := range opts {
		o(options)
	}
}

func defaultJsonEncoderFactory() transport.EncoderFactory {
	return func(w io.Writer) transport.Encoder {
		d := json.NewEncoder(w)
		return func(v interface{}) error {
			return d.Encode(v)
		}
	}
}

func defaultJsonDecoderFactory() transport.DecoderFactory {
	return func(r io.Reader) transport.Decoder {
		d := json.NewDecoder(r)
		return func(v interface{}) error {
			return d.Decode(v)
		}
	}
}
