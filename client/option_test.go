package client

import (
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/ekomobile/dadata/v2/client/transport"
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

func TestWithEncoderFactory(t *testing.T) {
	type args struct {
		c transport.EncoderFactory
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestWithEncoderFactory",
			args: args{
				c: func(w io.Writer) transport.Encoder {
					return func(v interface{}) error {
						return errors.New("c164a8d0-64b6-4374-a4b4-4036fbee504b")
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &clientOptions{}
			WithEncoderFactory(tt.args.c)(opts)

			assert.True(t, tt.args.c(nil)(nil).Error() == opts.encoderFactory(nil)(nil).Error())
		})
	}
}

func TestWithDecoderFactory(t *testing.T) {
	type args struct {
		c transport.DecoderFactory
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestWithDecoderFactory",
			args: args{
				c: func(r io.Reader) transport.Decoder {
					return func(v interface{}) error {
						return errors.New("b02ef946-15c8-40e0-b94d-efc3b26a8f75")
					}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &clientOptions{}
			WithDecoderFactory(tt.args.c)(opts)

			assert.True(t, tt.args.c(nil)(nil).Error() == opts.decoderFactory(nil)(nil).Error())
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
