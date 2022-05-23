package suggest

import (
	"context"
	"net/url"
)

type (
	// Requester provides transport level API calls.
	requester interface {
		// Get makes a GET API call. Assumes sending params in a request query string.
		Get(ctx context.Context, apiMethod string, params url.Values, result interface{}) error
		// Post makes a POST API call. Assumes sending json-encoded params in a request body.
		Post(ctx context.Context, apiMethod string, params, result interface{}) error
	}

	// ApiImp provides suggestion API.
	ApiImp struct {
		Client requester
	}

	Api interface {
		Address(context.Context, *RequestParams) ([]*AddressSuggestion, error)
		Bank(context.Context, *RequestParams) ([]*BankSuggestion, error)
		Email(context.Context, *RequestParams) ([]*EmailSuggestion, error)
		Country(context.Context, *RequestParams) ([]*CountrySuggestion, error)
		FMSUnit(context.Context, *RequestParams) ([]*FMSUnitSuggestion, error)
		AddressByID(context.Context, string) ([]*AddressSuggestion, error)
		CountryByID(context.Context, string) ([]*CountrySuggestion, error)
		GeoIP(context.Context, string) (*GeoIPResponse, error)
		Party(context.Context, *RequestParams) ([]*PartySuggestion, error)
		PartyByID(context.Context, *PartyByIDParams) ([]*PartySuggestion, error)
		Name(context.Context, *RequestParams, ...NameOption) ([]*NameSuggestion, error)
	}
)

// Address try to return suggest addresses by params
func (a *ApiImp) Address(ctx context.Context, params *RequestParams) (ret []*AddressSuggestion, err error) {
	var result = &AddressResponse{}
	err = a.Client.Post(ctx, "suggest/address", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Bank try to return suggest banks by params
func (a *ApiImp) Bank(ctx context.Context, params *RequestParams) (ret []*BankSuggestion, err error) {
	var result = &BankResponse{}
	err = a.Client.Post(ctx, "suggest/bank", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Email try to return suggest emails by params
func (a *ApiImp) Email(ctx context.Context, params *RequestParams) (ret []*EmailSuggestion, err error) {
	var result = &EmailResponse{}
	err = a.Client.Post(ctx, "suggest/email", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// Country try to return suggest countries by params
func (a *ApiImp) Country(ctx context.Context, params *RequestParams) (ret []*CountrySuggestion, err error) {
	var result = &CountryResponse{}
	err = a.Client.Post(ctx, "suggest/country", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// FMSUnit try to return suggest FMS unit by params
func (a *ApiImp) FMSUnit(ctx context.Context, params *RequestParams) (ret []*FMSUnitSuggestion, err error) {
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
func (a *ApiImp) AddressByID(ctx context.Context, id string) (addresses []*AddressSuggestion, err error) {
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
func (a *ApiImp) CountryByID(ctx context.Context, id string) (addresses []*CountrySuggestion, err error) {
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
func (a *ApiImp) GeoIP(ctx context.Context, ip string) (result *GeoIPResponse, err error) {
	result = &GeoIPResponse{}
	params := url.Values{
		"ip": []string{ip},
	}

	err = a.Client.Get(ctx, "iplocate/address", params, result)
	return
}
