package suggest

import (
	"context"
	"net/url"
)

type (
	// Requester provides transport level API calls.
	Requester interface {
		// Get makes a GET API call. Assumes sending params in a request query string.
		Get(ctx context.Context, apiMethod string, params url.Values, result interface{}) error
		// Post makes a POST API call. Assumes sending json-encoded params in a request body.
		Post(ctx context.Context, apiMethod string, params, result interface{}) error
	}

	// Api provides suggestion API.
	Api struct {
		Client Requester
	}
)

// Address try to return suggest addresses by params
func (a *Api) Address(ctx context.Context, params *RequestParams) (ret []*AddressSuggestion, err error) {
	var result = &AddressResponse{}
	err = a.Client.Post(ctx, "suggest/address", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Bank try to return suggest banks by params
func (a *Api) Bank(ctx context.Context, params *RequestParams) (ret []*BankSuggestion, err error) {
	var result = &BankResponse{}
	err = a.Client.Post(ctx, "suggest/bank", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Email try to return suggest emails by params
func (a *Api) Email(ctx context.Context, params *RequestParams) (ret []*EmailSuggestion, err error) {
	var result = &EmailResponse{}
	err = a.Client.Post(ctx, "suggest/email", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Country try to return suggest countries by params
func (a *Api) Country(ctx context.Context, params *RequestParams) (ret []*CountrySuggestion, err error) {
	var result = &CountryResponse{}
	err = a.Client.Post(ctx, "suggest/country", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// FMSUnit try to return suggest FMS unit by params
func (a *Api) FMSUnit(ctx context.Context, params *RequestParams) (ret []*FMSUnitSuggestion, err error) {
	var result = &FMSUnitResponse{}
	err = a.Client.Post(ctx, "suggest/fms_unit", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// AddressByID find addresses by Fias or Kladr
// see full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=312016944
func (a *Api) AddressByID(ctx context.Context, id string) (addresses []*AddressSuggestion, err error) {
	var result = &AddressResponse{}
	var req = &RequestParams{Query: id}

	err = a.Client.Post(ctx, "findById/address", req, result)
	if err != nil {
		return
	}
	addresses = result.Suggestions

	return
}

// CountryByID find countries by ID
func (a *Api) CountryByID(ctx context.Context, id string) (addresses []*CountrySuggestion, err error) {
	var result = &CountryResponse{}
	var req = &RequestParams{Query: id}

	err = a.Client.Post(ctx, "findById/country", req, result)
	if err != nil {
		return
	}
	addresses = result.Suggestions

	return
}

// GeoIP try to find address by IP
// see documentation on: https://dadata.ru/api/detect_address_by_ip/
// ip string representation of ip-address (example 10.12.44.23)
// if ip=="" then dadata try to get ip-address from X-Forwarded-For header
func (a *Api) GeoIP(ctx context.Context, ip string) (result *GeoIPResponse, err error) {
	result = &GeoIPResponse{}
	params := url.Values{
		"ip": []string{ip},
	}

	err = a.Client.Get(ctx, "iplocate/address", params, result)
	return
}
