package dadata

// GeoIPResponse response for GeoIP
type GeoIPResponse struct {
	Location *ResponseAddress `json:"location"`
}
