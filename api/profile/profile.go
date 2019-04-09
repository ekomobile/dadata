package profile

import (
	"context"
	"net/url"
)

type (
	// Requester provides transport level API calls.
	Requester interface {
		// Get makes a GET API call. Assumes sending params in a request query string.
		Get(ctx context.Context, apiMethod string, params url.Values, result interface{}) error
	}

	// Api provides profile related API.
	Api struct {
		Client Requester
	}
)

// Balance return daily statistics
// see documentation https://dadata.ru/api/stat/
func (a *Api) Balance(ctx context.Context) (result *BalanceResponse, err error) {
	result = &BalanceResponse{}
	err = a.Client.Get(ctx, "profile/balance", nil, result)
	return
}
