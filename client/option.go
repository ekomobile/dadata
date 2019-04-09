package client

import (
	"net/http"
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

func applyOptions(options *clientOptions, opts ...Option) {
	for _, o := range opts {
		o(options)
	}
}
