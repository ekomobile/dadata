package suggest

import (
	"github.com/ekomobile/dadata/v2/api/model"
)

type (
	// GeolocateParams Request struct
	// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=808583277
	GeolocateParams struct {
		Lat          string `json:"lat"`                     // geographic latitude
		Lon          string `json:"lon"`                     // geographic longitude
		Count        string `json:"count,omitempty"`         // number of results (max 20)
		RadiusMeters string `json:"radius_meters,omitempty"` // search radius in metres (max. 1000)
		Language     string `json:"language,omitempty"`      // in which language to return the result (ru / en)
	}

	// GeolocateSuggestion api response for address
	GeolocateSuggestion struct {
		Value             string         `json:"value"`
		UnrestrictedValue string         `json:"unrestricted_value"`
		Data              *model.Address `json:"data"`
	}

	// GeolocateResponse result slice for address suggestions
	GeolocateResponse struct {
		Suggestions []*GeolocateSuggestion `json:"suggestions"`
	}
)
