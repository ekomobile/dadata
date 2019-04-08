package dadata

import (
	"context"
	"fmt"
)

// AddressByID find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (c *Client) AddressByID(id string) (*ResponseAddress, error) {
	return c.AddressByIDWithCtx(context.Background(), id)
}

// AddressByIDWithCtx find address by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (c *Client) AddressByIDWithCtx(ctx context.Context, id string) (address *ResponseAddress, err error) {
	var result []ResponseAddress
	if result, err = c.AddressesByIDWithCtx(ctx, id); err != nil {
		return
	}
	address = &result[0]
	return
}

// AddressesByID find addresses by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (c *Client) AddressesByID(id string) ([]ResponseAddress, error) {
	return c.AddressesByIDWithCtx(context.Background(), id)
}

// AddressesByIDWithCtx find addresses by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (c *Client) AddressesByIDWithCtx(ctx context.Context, id string) (addresses []ResponseAddress, err error) {
	var result = &SuggestAddressResponse{}
	var req = SuggestRequestParams{Query: id}

	if err = c.sendRequestToURL(ctx, "POST", c.options.baseSuggestURL+"findById/address", req, result); err != nil {
		return
	}
	if len(result.Suggestions) == 0 {
		err = fmt.Errorf("dadata.AddressByID: cannot detect address by id %s", id)
		return
	}
	addresses = result.Suggestions

	return
}

// CountryByID find country by ID
func (c *Client) CountryByID(id string) (*ResponseCountry, error) {
	return c.CountryByIDWithCtx(context.Background(), id)
}

// CountryByIDWithCtx find country by ID
func (c *Client) CountryByIDWithCtx(ctx context.Context, id string) (country *ResponseCountry, err error) {
	var result []ResponseCountry
	if result, err = c.CountriesByIDWithCtx(ctx, id); err != nil {
		return
	}
	country = &result[0]
	return
}

// CountriesByID find countries by ID
func (c *Client) CountriesByID(id string) ([]ResponseCountry, error) {
	return c.CountriesByIDWithCtx(context.Background(), id)
}

// CountriesByIDWithCtx find countries by ID
func (c *Client) CountriesByIDWithCtx(ctx context.Context, id string) (addresses []ResponseCountry, err error) {
	var result = &SuggestCountryResponse{}
	var req = SuggestRequestParams{Query: id}

	if err = c.sendRequestToURL(ctx, "POST", c.options.baseSuggestURL+"findById/country", req, result); err != nil {
		return
	}
	if len(result.Suggestions) == 0 {
		err = fmt.Errorf("dadata.CountryByID: cannot detect country by id %s", id)
		return
	}
	addresses = result.Suggestions

	return
}
