package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHttpClient(t *testing.T) {
	tests := []struct {
		name string
		cli  *http.Client
	}{
		{
			name: "WithHttpClient",
			cli:  http.DefaultClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &clientOptions{}
			WithHttpClient(tt.cli)(opts)

			assert.Equal(t, tt.cli, opts.httpClient)
		})
	}
}

func TestWithCredentialProvider(t *testing.T) {
	type args struct {
		c CredentialProvider
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "WithCredentialProvider", args: args{c: &Credentials{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &clientOptions{}
			WithCredentialProvider(tt.args.c)(opts)

			assert.Equal(t, tt.args.c, opts.credentialProvider)
		})
	}
}

func Test_applyOptions(t *testing.T) {
	cp := &Credentials{}

	type args struct {
		options *clientOptions
		opts    []Option
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "applyOptions", args: args{
			opts: []Option{
				WithHttpClient(http.DefaultClient),
				WithCredentialProvider(cp),
			},
			options: &clientOptions{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			applyOptions(tt.args.options, tt.args.opts...)

			assert.Equal(t, cp, tt.args.options.credentialProvider)
			assert.Equal(t, http.DefaultClient, tt.args.options.httpClient)
		})
	}
}
