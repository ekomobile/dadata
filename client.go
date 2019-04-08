// Golang client library for DaData.ru (https://dadata.ru/).

// Package dadata implemented cleaning (https://dadata.ru/api/clean/) and suggesting (https://dadata.ru/api/suggest/)
package dadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	defaultBaseURL        = "https://dadata.ru/api/v2/"
	defaultBaseSuggestURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/"
)

type (
	Client struct {
		options clientOptions
	}

	CredentialProvider interface {
		ApiKey() string
		SecretKey() string
	}

	Credentials struct {
		ApiKeyValue    string
		SecretKeyValue string
	}

	EnvironmentCredentials struct {
		ApiKeyName    string
		SecretKeyName string
	}

	clientOptions struct {
		httpClient         *http.Client
		credentialProvider CredentialProvider
		baseURL            string
		baseSuggestURL     string
	}

	ClientOption func(opts *clientOptions)
)

// NewClient Create new client of DaData.
// Api and secret keys see on profile page (https://dadata.ru/profile/).
// By default client uses `DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables.
func NewClient(opts ...ClientOption) *Client {
	options := clientOptions{
		httpClient: http.DefaultClient,
		credentialProvider: &EnvironmentCredentials{
			ApiKeyName:    "DADATA_API_KEY",
			SecretKeyName: "DADATA_SECRET_KEY",
		},
		baseURL:        defaultBaseURL,
		baseSuggestURL: defaultBaseSuggestURL,
	}

	applyOptions(&options, opts...)

	return &Client{
		options: options,
	}
}

func (c *Client) sendRequestToURL(ctx context.Context, method, url string, source interface{}, result interface{}) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("sendRequestToURL: ctx.Err return err=%v", err)
	}

	buffer := &bytes.Buffer{}

	if err := json.NewEncoder(buffer).Encode(source); err != nil {
		return fmt.Errorf("sendRequestToURL: json.Encode return err = %v", err)
	}

	request, err := http.NewRequest(method, url, buffer)

	if err != nil {
		return fmt.Errorf("sendRequestToURL: http.NewRequest return err = %v", err)
	}

	request = request.WithContext(ctx)

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", c.options.credentialProvider.ApiKey()))
	request.Header.Add("X-Secret", c.options.credentialProvider.SecretKey())
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := c.options.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("sendRequestToURL: httpClient.Do return err = %v", err)
	}

	defer response.Body.Close()

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("sendRequestToURL: Request error %v", response.Status)
	}

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("sendRequestToURL: json.Decode return err = %v", err)
	}

	return nil
}

// sendRequest
func (c *Client) sendRequest(ctx context.Context, lastURLPart string, source interface{}, result interface{}) error {
	return c.sendRequestToURL(ctx, "POST", c.options.baseURL+lastURLPart, source, result)
}

func WithHttpClient(c *http.Client) ClientOption {
	return func(opts *clientOptions) {
		opts.httpClient = c
	}
}

// WithCredentialProvider sets credential provider.
func WithCredentialProvider(c CredentialProvider) ClientOption {
	return func(opts *clientOptions) {
		opts.credentialProvider = c
	}
}

func WithBaseURL(url string) ClientOption {
	return func(opts *clientOptions) {
		opts.baseURL = url
	}
}

func WithBaseSuggestURL(url string) ClientOption {
	return func(opts *clientOptions) {
		opts.baseSuggestURL = url
	}
}

func applyOptions(options *clientOptions, opts ...ClientOption) {
	options = &clientOptions{}
	for _, o := range opts {
		o(options)
	}
}

func (c *Credentials) ApiKey() string {
	return c.ApiKeyValue
}

func (c *Credentials) SecretKey() string {
	return c.SecretKeyValue
}

func (c *EnvironmentCredentials) ApiKey() string {
	return os.Getenv(c.ApiKeyName)
}

func (c *EnvironmentCredentials) SecretKey() string {
	return os.Getenv(c.SecretKeyName)
}
