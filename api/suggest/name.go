package suggest

import (
	"context"

	"github.com/ekomobile/dadata/v2/api/model"
)

// Name gender values
// https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669115
const (
	NameGenderUnknown = "UNKNOWN" // не удалось однозначно определить
	NameGenderMale    = "MALE"
	NameGenderFemale  = "FEMALE"
)

// Name parts
// https://dadata.ru/api/suggest/name/
// https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669115
const (
	NamePartSurname    = "SURNAME"
	NamePartName       = "NAME"
	NamePartPatronymic = "PATRONYMIC"
)

type (
	NameParams struct {
		Query  string   `json:"query"`
		Count  int      `json:"count"`
		Parts  []string `json:"parts"`
		Gender string   `json:"gender"`
	}
	NameOption func(params *NameParams)

	// NameResponse result slice for name suggestions
	NameResponse struct {
		Suggestions []*NameSuggestion `json:"suggestions"`
	}

	// NameSuggestion api response for name
	NameSuggestion struct {
		Value             string      `json:"value"`
		UnrestrictedValue string      `json:"unrestricted_value"`
		Data              *model.Name `json:"data"`
	}
)

// WithNameParts adds `parts` parameter to suggest/fio request
// https://dadata.ru/api/suggest/name/
// https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669115
func WithNameParts(parts ...string) NameOption {
	return func(params *NameParams) {
		params.Parts = parts
	}
}

// WithNameGender adds `gender` parameter to suggest/fio request
// https://dadata.ru/api/suggest/name/
// https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669115
func WithNameGender(gender string) NameOption {
	return func(params *NameParams) {
		params.Gender = gender
	}
}

func (p *NameParams) applyOption(opts ...NameOption) {
	for _, o := range opts {
		o(p)
	}
}

// Name try to return suggest names by params
func (a *Api) Name(ctx context.Context, requestParams *RequestParams, opts ...NameOption) (ret []*NameSuggestion, err error) {
	var result = &NameResponse{}

	params := NameParams{
		Query: requestParams.Query,
		Count: requestParams.Count,
	}
	params.applyOption(opts...)

	err = a.Client.Post(ctx, "suggest/fio", &params, result)
	if err != nil {
		return
	}
	ret = result.Suggestions
	return
}
