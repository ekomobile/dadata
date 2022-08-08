package suggest

import (
	"github.com/ekomobile/dadata/v2/api/model"
)

type (
	// RequestParamsLocation constraints for suggestion
	// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669108
	RequestParamsLocation struct {
		FiasID               string `json:"fias_id,omitempty"`
		KladrID              string `json:"kladr_id,omitempty"`
		Region               string `json:"region,omitempty"`
		RegionFiasID         string `json:"region_fias_id,omitempty"`
		RegionKladrID        string `json:"region_kladr_id,omitempty"`
		RegionTypeFull       string `json:"region_type_full,omitempty"`
		City                 string `json:"city,omitempty"`
		CityFiasID           string `json:"city_fias_id,omitempty"` // search only in this area
		CityKladrID          string `json:"city_kladr_id,omitempty"`
		CityTypeFull         string `json:"city_type_full,omitempty"`
		CityDistrictTypeFull string `json:"city_district_type_full,omitempty"`
		Settlement           string `json:"settlement,omitempty"`
		SettlementFiasID     string `json:"settlement_fias_id,omitempty"`
		SettlementKladrID    string `json:"settlement_kladr_id,omitempty"`
		SettlementTypeFull   string `json:"settlement_type_full,omitempty"`
		Street               string `json:"street,omitempty"`
		StreetFiasID         string `json:"street_fias_id,omitempty"`
		StreetKladrID        string `json:"street_kladr_id,omitempty"`
		StreetTypeFull       string `json:"street_type_full,omitempty"`
		AreaTypeFull         string `json:"area_type_full,omitempty"`
		Country              string `json:"country,omitempty"`
		CountryISOCode       string `json:"country_iso_code,omitempty"`
	}

	// Bound for granular sugestion
	// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=222888017
	Bound struct {
		Value model.BoundValue `json:"value"`
	}

	// RequestParams Request struct
	RequestParams struct {
		Query         string                   `json:"query"` // user input for suggestion
		Count         int                      `json:"count"` // ligmit for results
		Locations     []*RequestParamsLocation `json:"locations"`
		RestrictValue bool                     `json:"restrict_value"` // don't show restricts (region) on results

		FromBound *Bound `json:"from_bound"`
		ToBound   *Bound `json:"to_bound"`
	}

	// AddressResponse result slice for address suggestions
	AddressResponse struct {
		Suggestions []*AddressSuggestion `json:"suggestions"`
	}

	// BankResponse result slice for bank suggestions
	BankResponse struct {
		Suggestions []*BankSuggestion `json:"suggestions"`
	}

	// PartyResponse result slice for party suggestions
	PartyResponse struct {
		Suggestions []*PartySuggestion `json:"suggestions"`
	}

	// EmailResponse result slice for email suggestions
	EmailResponse struct {
		Suggestions []*EmailSuggestion `json:"suggestions"`
	}

	// CountryResponse result slice for country suggestions
	CountryResponse struct {
		Suggestions []*CountrySuggestion `json:"suggestions"`
	}

	// FMSUnitResponse result slice for FMS unit suggestions
	FMSUnitResponse struct {
		Suggestions []*FMSUnitSuggestion `json:"suggestions"`
	}

	// GeoIPResponse response for GeoIP
	GeoIPResponse struct {
		Location *AddressSuggestion `json:"location"`
	}

	// BankSuggestion api response for bank
	BankSuggestion struct {
		Value             string      `json:"value"`
		UnrestrictedValue string      `json:"unrestricted_value"`
		Data              *model.Bank `json:"data"`
	}

	// AddressSuggestion api response for address
	AddressSuggestion struct {
		Value             string         `json:"value"`
		UnrestrictedValue string         `json:"unrestricted_value"`
		Data              *model.Address `json:"data"`
	}

	// PartySuggestion api response for party
	// https://confluence.hflabs.ru/pages/viewpage.action?pageId=204669122
	PartySuggestion struct {
		Value             string       `json:"value"`
		UnrestrictedValue string       `json:"unrestricted_value"`
		Data              *model.Party `json:"data"`
	}

	// CountrySuggestion api response for country
	CountrySuggestion struct {
		Value string         `json:"value"`
		Data  *model.Country `json:"data"`
	}

	// EmailSuggestion api response for email
	EmailSuggestion struct {
		Value             string       `json:"value"`
		UnrestrictedValue string       `json:"unrestricted_value"`
		Data              *model.Email `json:"data"`
	}

	// FMSUnitSuggestion is a FMS unit suggestion
	// https://dadata.ru/api/suggest/fms_unit/
	FMSUnitSuggestion struct {
		Value             string         `json:"value"`
		UnrestrictedValue string         `json:"unrestricted_value"`
		Data              *model.FMSUnit `json:"data"`
	}
)
