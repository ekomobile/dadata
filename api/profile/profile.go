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

	// ApiImp provides profile related API.
	ApiImp struct {
		Client Requester
	}

	Api interface {
		Balance(context.Context) (*BalanceResponse, error)
	}
)

// Balance return daily statistics
// see documentation https://dadata.ru/api/stat/
func (a *ApiImp) Balance(ctx context.Context) (result *BalanceResponse, err error) {
	result = &BalanceResponse{}
	err = a.Client.Get(ctx, "profile/balance", nil, result)
	return
}
