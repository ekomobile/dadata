package dadata

import "context"

func (c *Client) sendSuggestRequest(ctx context.Context, lastURLPart string, requestParams SuggestRequestParams, result interface{}) error {
	return c.sendRequest(ctx, "suggest/"+lastURLPart, requestParams, result)
}

// SuggestAddresses try to return suggest addresses by requestParams
func (c *Client) SuggestAddresses(requestParams SuggestRequestParams) ([]ResponseAddress, error) {
	return c.SuggestAddressesWithCtx(context.Background(), requestParams)
}

// SuggestAddressesWithCtx try to return suggest addresses by requestParams
func (c *Client) SuggestAddressesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseAddress, err error) {
	var result = &SuggestAddressResponse{}
	if err = c.sendSuggestRequest(ctx, "address", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestNames try to return suggest names by requestParams
func (c *Client) SuggestNames(requestParams SuggestRequestParams) ([]ResponseName, error) {
	return c.SuggestNamesWithCtx(context.Background(), requestParams)
}

// SuggestNamesWithCtx try to return suggest names by requestParams
func (c *Client) SuggestNamesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseName, err error) {
	var result = &SuggestNameResponse{}
	if err = c.sendSuggestRequest(ctx, "fio", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestBanks try to return suggest banks by requestParams
func (c *Client) SuggestBanks(requestParams SuggestRequestParams) ([]ResponseBank, error) {
	return c.SuggestBanksWithCtx(context.Background(), requestParams)
}

// SuggestBanksWithCtx try to return suggest banks by requestParams
func (c *Client) SuggestBanksWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseBank, err error) {
	var result = &SuggestBankResponse{}
	if err = c.sendSuggestRequest(ctx, "bank", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestParties try to return suggest parties by requestParams
func (c *Client) SuggestParties(requestParams SuggestRequestParams) ([]ResponseParty, error) {
	return c.SuggestPartiesWithCtx(context.Background(), requestParams)
}

// SuggestPartiesWithCtx try to return suggest parties by requestParams
func (c *Client) SuggestPartiesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseParty, err error) {
	var result = &SuggestPartyResponse{}
	if err = c.sendSuggestRequest(ctx, "party", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestEmails try to return suggest emails by requestParams
func (c *Client) SuggestEmails(requestParams SuggestRequestParams) ([]ResponseEmail, error) {
	return c.SuggestEmailsWithCtx(context.Background(), requestParams)
}

// SuggestEmailsWithCtx try to return suggest emails by requestParams
func (c *Client) SuggestEmailsWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseEmail, err error) {
	var result = &SuggestEmailResponse{}
	if err = c.sendSuggestRequest(ctx, "email", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}

// SuggestCountries try to return suggest countries by requestParams
func (c *Client) SuggestCountries(requestParams SuggestRequestParams) ([]ResponseCountry, error) {
	return c.SuggestCountriesWithCtx(context.Background(), requestParams)
}

// SuggestCountriesWithCtx try to return suggest countries by requestParams
func (c *Client) SuggestCountriesWithCtx(ctx context.Context, requestParams SuggestRequestParams) (ret []ResponseCountry, err error) {
	var result = &SuggestCountryResponse{}
	if err = c.sendSuggestRequest(ctx, "country", requestParams, result); err != nil {
		return
	}
	ret = result.Suggestions
	return
}
