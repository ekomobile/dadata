package geolocate

import (
	"context"
)

type (
	// Requester provides transport level API calls.
	Requester interface {
		// Post makes a POST API call. Assumes sending json-encoded params in a request body.
		Post(ctx context.Context, apiMethod string, params, result interface{}) error
	}

	// Api provides suggestion API.
	Api struct {
		Client Requester
	}
)

// AddressByCoordinates try to return suggest addresses by params
// see documentation on:https://dadata.ru/api/geolocate/
func (a *Api) AddressByCoordinates(ctx context.Context, params *RequestParams) (ret []*AddressSuggestion, err error) {
	var result = &AddressResponse{}
	err = a.Client.Post(ctx, "geolocate/address", params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}
