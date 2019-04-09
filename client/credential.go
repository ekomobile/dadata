package client

import (
	"os"
)

type (
	// Credentials provides constant credential values.
	Credentials struct {
		ApiKeyValue    string
		SecretKeyValue string
	}

	// EnvironmentCredentials provides credentials from environment variables.
	EnvironmentCredentials struct {
		ApiKeyName    string
		SecretKeyName string
	}
)

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
