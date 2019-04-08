package dadata

import (
	"context"
	"time"
)

// DailyStat return daily statistics
// see documentation https://dadata.ru/api/stat/
func (c *Client) DailyStat(date time.Time) (*StatResponse, error) {
	return c.DailyStatWithCtx(context.Background(), date)
}

// DailyStatWithCtx return daily statistics
// see documentation https://dadata.ru/api/stat/
func (c *Client) DailyStatWithCtx(ctx context.Context, date time.Time) (result *StatResponse, err error) {
	var dateStr string

	result, dateStr = &StatResponse{}, date.Format("2006-01-02")
	if err = c.sendRequestToURL(ctx, "GET", c.options.baseURL+"stat/daily?date="+dateStr, nil, result); err != nil {
		result = nil
	}

	return
}
