package dadata

import (
	"context"
	"fmt"
	"time"
)

func ExampleClient_SuggestAddresses() {
	daData := NewClient()

	addresses, err := daData.SuggestAddresses(SuggestRequestParams{Query: "Преснен", Count: 2})
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
		fmt.Println(address.Data.Street)
		fmt.Println(address.Data.FiasLevel)
	}
}

func ExampleClient_SuggestAddresses_granular() {
	daData := NewClient()

	var req SuggestRequestParams

	req.Query = "лен"

	req.Locations = append(req.Locations, SuggestRequestParamsLocation{
		RegionFiasID: "df3d7359-afa9-4aaa-8ff9-197e73906b1c",
		CityFiasID:   "e9e684ce-7d60-4480-ba14-ca6da658188b",
	})

	req.FromBound = SuggestBound{SuggestBoundStreet}
	req.ToBound = SuggestBound{SuggestBoundStreet}

	req.RestrictValue = true
	req.Count = 2

	addresses, err := daData.SuggestAddresses(req)
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
		fmt.Println(address.Data.Street)
	}
}

func ExampleClient_SuggestAddressesWithCtx() {
	daData := NewClient()

	var req SuggestRequestParams

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	req.Query = "лен"

	req.Locations = append(req.Locations, SuggestRequestParamsLocation{
		RegionFiasID: "df3d7359-afa9-4aaa-8ff9-197e73906b1c",
		CityFiasID:   "e9e684ce-7d60-4480-ba14-ca6da658188b",
	})

	req.FromBound = SuggestBound{SuggestBoundStreet}
	req.ToBound = SuggestBound{SuggestBoundStreet}

	req.RestrictValue = true
	req.Count = 2

	addresses, err := daData.SuggestAddressesWithCtx(ctx, req)
	if nil != err {
		fmt.Println(err)
	}

	for _, address := range addresses {
		fmt.Println(address.UnrestrictedValue)
	}

	cancel()
	// if ctx is exited (by cancel or timeout) we must catch err
	_, err = daData.SuggestAddressesWithCtx(ctx, req)
	fmt.Println(err)
}

func ExampleClient_SuggestBanks() {
	daData := NewClient()

	banks, err := daData.SuggestBanks(SuggestRequestParams{Query: "Кредитный", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, bank := range banks {
		fmt.Println(bank.Data.Name.Full)
		fmt.Println(bank.Data.Bic)
	}
}

func ExampleClient_SuggestParties() {
	daData := NewClient()

	parties, err := daData.SuggestParties(SuggestRequestParams{Query: "Агрохолд", Count: 3})
	if nil != err {
		fmt.Println(err)
	}

	for _, party := range parties {
		fmt.Println(party.Data.Name.Full)
		fmt.Println(party.Data.Ogrn)
	}
}
