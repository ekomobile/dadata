package dadata

import (
	"context"
)

// ProfileBalance return daily statistics
// see documentation https://dadata.ru/api/stat/
func (c *Client) ProfileBalance() (*BalanceResponse, error) {
	return c.ProfileBalanceWithCtx(context.Background())
}

// ProfileBalanceWithCtx return daily statistics
// see documentation https://dadata.ru/api/stat/
func (c *Client) ProfileBalanceWithCtx(ctx context.Context) (result *BalanceResponse, err error) {
	result = new(BalanceResponse)
	if err = c.sendRequestToURL(ctx, "GET", c.options.baseURL+"profile/balance", nil, result); err != nil {
		result = nil
	}
	return
}
