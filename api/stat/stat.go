package stat

import (
	"context"
	"net/url"
	"time"
)

type (
	// Requester provides transport level API calls.
	Requester interface {
		// Get makes a GET API call. Assumes sending params in a request query string.
		Get(ctx context.Context, apiMethod string, params url.Values, result interface{}) error
	}

	// Api provides statistic API.
	Api struct {
		Client Requester
	}
)

// Daily return daily statistics
// see documentation https://dadata.ru/api/stat/
func (a *Api) Daily(ctx context.Context, date time.Time) (result *Stat, err error) {
	result = &Stat{}
	params := url.Values{
		"date": []string{date.Format("2006-01-02")},
	}

	err = a.Client.Get(ctx, "stat/daily", params, result)
	return
}
