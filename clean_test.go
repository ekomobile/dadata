package dadata

import (
	"fmt"
)

func ExampleClient_CleanAddresses() {
	// By default client uses `DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables.
	daData := NewClient()
	// Or credentials may be passed as client option.
	// daData := NewClient(WithCredentialProvider("API_KEY", "SECRET_KEY"))

	addresses, err := daData.CleanAddresses("ул.Правды 26", "пер.Расковой 5")

	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.StreetTypeFull)
		fmt.Println(address.Street)
		fmt.Println(address.House)
	}
}

func ExampleClient_CleanNames() {
	daData := NewClient()

	names, err := daData.CleanNames("Алексей Иванов", "Иван Алексеев")
	if nil != err {
		fmt.Println(err)
	}

	for _, name := range names {
		fmt.Println(name.Surname)
		fmt.Println(name.Name)
	}
}
