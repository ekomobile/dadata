package dadata

import (
	"fmt"
)

func ExampleDaData_GeoIP() {
	daData := NewClient()

	geoIPResponse, err := daData.GeoIP("83.220.54.223")
	if nil != err {
		fmt.Println(err)
		return
	}
	if geoIPResponse.Location == nil {
		fmt.Println("empty result from GeoIP")
		return
	}
	address := geoIPResponse.Location.Data
	fmt.Println(address.Country)
	fmt.Println(address.City)
	fmt.Printf("see on https://www.google.com/maps/@%s,%sf,14z\n", address.GeoLat, address.GeoLon)
}
