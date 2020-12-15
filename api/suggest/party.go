package suggest

import (
	"context"
)

type (
	PartyByIDParams struct {
		Query      string           `json:"query"`
		Count      *int             `json:"count,omitempty"`
		KPP        *string          `json:"kpp,omitempty"`
		Type       *PartyType       `json:"type,omitempty"`
		BranchType *PartyBranchType `json:"branch_type,omitempty"`
	}

	PartyType       string
	PartyBranchType string
)

const (
	PartyTypeLegal      PartyType = "LEGAL"
	PartyTypeIndividual PartyType = "INDIVIDUAL"

	PartyBranchTypeMain   PartyBranchType = "MAIN"
	PartyBranchTypeBranch PartyBranchType = "BRANCH"
)

func NewPartyByIDParams(query string) *PartyByIDParams {
	return &PartyByIDParams{
		Query: query,
	}
}

func (o *PartyByIDParams) SetQuery(query string) *PartyByIDParams {
	o.Query = query
	return o
}

func (o *PartyByIDParams) SetCount(count int) *PartyByIDParams {
	o.Count = &count
	return o
}

func (o *PartyByIDParams) SetKPP(kpp string) *PartyByIDParams {
	o.KPP = &kpp
	return o
}

func (o *PartyByIDParams) SetType(t PartyType) *PartyByIDParams {
	o.Type = &t
	return o
}

func (o *PartyByIDParams) SetBranchType(t PartyBranchType) *PartyByIDParams {
	o.BranchType = &t
	return o
}

// Party try to return suggest parties by params
func (a *Api) Party(ctx context.Context, params *RequestParams) (ret []*PartySuggestion, err error) {
	var result = &PartyResponse{}

	err = a.Client.Post(ctx, "suggest/party", params, result)
	if err != nil {
		return
	}

	ret = result.Suggestions

	return
}

// Party find parties by ID
// https://dadata.ru/api/find-party/
func (a *Api) PartyByID(ctx context.Context, params *PartyByIDParams) (ret []*PartySuggestion, err error) {
	var result = &PartyResponse{}

	err = a.Client.Post(ctx, "findById/party", params, result)
	if err != nil {
		return
	}

	ret = result.Suggestions

	return
}
