package geolocate

import "github.com/ekomobile/dadata/v2/api/model"

type (
	// RequestParams Request struct
	// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=808583277
	RequestParams struct {
		Lat          string `json:"lat"`                     // geographic latitude
		Lon          string `json:"lon"`                     // geographic longitude
		Count        string `json:"count,omitempty"`         // number of results (max 20)
		RadiusMeters string `json:"radius_meters,omitempty"` // search radius in metres (max. 1000)
		Language     string `json:"language,omitempty"`      // in which language to return the result (ru / en)
	}

	// AddressGeolocateSuggestion api response for address
	AddressGeolocateSuggestion struct {
		Value             string         `json:"value"`
		UnrestrictedValue string         `json:"unrestricted_value"`
		Data              *model.Address `json:"data"`
	}

	// AddressGeolocateResponse result slice for address suggestions
	AddressGeolocateResponse struct {
		Suggestions []*AddressGeolocateSuggestion `json:"suggestions"`
	}
)
