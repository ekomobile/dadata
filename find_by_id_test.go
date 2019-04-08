package dadata

import (
	"fmt"
)

func ExampleClient_AddressByID() {
	daData := NewClient()
	addr, err := daData.AddressByID("6300000100000")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("City: %s,\nFiasID: %s,\nKladr: %s\n", addr.Data.City, addr.Data.FiasID, addr.Data.KladrID)
}
