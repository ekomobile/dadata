package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	// Client is a transport for API calls.
	Client struct {
		options clientOptions
	}

	// CredentialProvider provides DaData API access keys.
	CredentialProvider interface {
		ApiKey() string
		SecretKey() string
	}

	clientOptions struct {
		httpClient         *http.Client
		credentialProvider CredentialProvider
		endpointURL        *url.URL
	}
)

// NewClient Create new client of DaData.
// Api and secret keys see on profile page (https://dadata.ru/profile/).
// By default client uses `DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables.
func NewClient(endpointURL *url.URL, opts ...Option) *Client {
	options := clientOptions{
		httpClient: http.DefaultClient,
		credentialProvider: &EnvironmentCredentials{
			ApiKeyName:    "DADATA_API_KEY",
			SecretKeyName: "DADATA_SECRET_KEY",
		},
		endpointURL: endpointURL,
	}

	applyOptions(&options, opts...)

	return &Client{
		options: options,
	}
}

func (c *Client) doRequest(ctx context.Context, method string, url *url.URL, body interface{}, result interface{}) (err error) {
	if err = ctx.Err(); err != nil {
		return fmt.Errorf("doRequest: ctx.Err return err=%v", err)
	}

	buffer := &bytes.Buffer{}

	if err = json.NewEncoder(buffer).Encode(body); err != nil {
		return fmt.Errorf("doRequest: json.Encode return err = %v", err)
	}

	request, err := http.NewRequest(method, url.String(), buffer)
	if err != nil {
		return fmt.Errorf("doRequest: http.NewRequest return err = %v", err)
	}

	request = request.WithContext(ctx)

	request.Header.Add("Authorization", fmt.Sprintf("Token %s", c.options.credentialProvider.ApiKey()))
	request.Header.Add("X-Secret", c.options.credentialProvider.SecretKey())
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := c.options.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("doRequest: httpClient.Do return err = %v", err)
	}

	defer response.Body.Close()

	if http.StatusOK != response.StatusCode {
		return fmt.Errorf("doRequest: Request error %v", response.Status)
	}

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("doRequest: json.Decode return err = %v", err)
	}

	return
}

func (c *Client) Post(ctx context.Context, apiMethod string, body, result interface{}) (err error) {
	requestURL, err := c.options.endpointURL.Parse(apiMethod)
	if err != nil {
		return
	}

	return c.doRequest(ctx, "POST", requestURL, body, result)
}

func (c *Client) Get(ctx context.Context, apiMethod string, params url.Values, result interface{}) (err error) {
	requestURL, err := c.options.endpointURL.Parse(apiMethod)
	if err != nil {
		return
	}

	if params != nil {
		q := requestURL.Query()
		for k, v := range params {
			q[k] = v
		}
		requestURL.RawQuery = q.Encode()
	}

	return c.doRequest(ctx, "GET", requestURL, nil, result)
}
