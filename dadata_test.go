// Live API tests.
// To run API test pass `DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables with your values.
package dadata

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ekomobile/dadata/v2/api/suggest"
)

type (
	ApiSuggestIntegrationTest struct {
		suite.Suite
	}

	ApiCleanIntegrationTest struct {
		suite.Suite
	}
)

func (s *ApiSuggestIntegrationTest) SetupSuite() {
	if _, ok := os.LookupEnv("DADATA_API_KEY"); !ok {
		s.Suite.T().Skip("no api keys in env")
	}
}

func (s *ApiCleanIntegrationTest) SetupSuite() {
	if _, ok := os.LookupEnv("DADATA_API_KEY"); !ok {
		s.Suite.T().Skip("no api keys in env")
	}
}

func TestSuggestApiIntegration(t *testing.T) {
	suite.Run(t, &ApiSuggestIntegrationTest{})
}

func TestCleanApiIntegration(t *testing.T) {
	suite.Run(t, &ApiCleanIntegrationTest{})
}

func (s *ApiSuggestIntegrationTest) TestAddress() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "ул Свободы",
	}
	res, err := api.Address(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestAddressWithLanguageParamRU() {
	api := NewSuggestApi()

	params := suggest.RequestParams{
		Query:    "ул Свободы",
		Language: "RU",
	}

	res, err := api.Address(context.Background(), &params)

	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestAddressWithLanguageParamEN() {
	api := NewSuggestApi()

	params := suggest.RequestParams{
		Query:    "ул Свободы",
		Language: "EN",
	}

	res, err := api.Address(context.Background(), &params)

	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestBank() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "сбербанк",
	}
	res, err := api.Bank(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestParty() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "сбербанк",
	}
	res, err := api.Party(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestPartyWithTypePositive() {
	api := NewSuggestApi()

	// positive: there are some legal entities "сбербанк"
	legalParty := suggest.PartyTypeLegal
	res, err := api.Party(context.Background(), &suggest.RequestParams{
		Query: "сбербанк",
		Type:  &legalParty,
	})

	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestPartyWithTypeNegative() {
	api := NewSuggestApi()

	// there are no one individual entrepreneur "сбербанк"
	individualParty := suggest.PartyTypeIndividual
	res, err := api.Party(context.Background(), &suggest.RequestParams{
		Query: "сбербанк",
		Type:  &individualParty,
	})

	s.NoError(err)
	s.Empty(res)
}

func (s *ApiSuggestIntegrationTest) TestPartyById() {
	api := NewSuggestApi()

	query := "7707083893"
	params := suggest.NewPartyByIDParams(query)

	res, err := api.PartyByID(context.Background(), params)

	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestPartyByIdWithKPP() {
	api := NewSuggestApi()

	query := "7707083893"
	testKPP := "773601001"

	params := suggest.NewPartyByIDParams(query).
		SetKPP(testKPP)

	res, err := api.PartyByID(context.Background(), params)

	s.NoError(err)
	s.NotEmpty(res)
	s.Equal(query, res[0].Data.Inn)
	s.Equal(testKPP, res[0].Data.Kpp)
}

func (s *ApiSuggestIntegrationTest) TestPartyByIdWithBranchType() {
	api := NewSuggestApi()

	query := "7707083893"

	params := suggest.NewPartyByIDParams(query).
		SetBranchType(suggest.PartyBranchTypeMain).
		SetType(suggest.PartyTypeLegal)

	res, err := api.PartyByID(context.Background(), params)

	s.NoError(err)
	s.NotEmpty(res)
	s.Equal(query, res[0].Data.Inn)
}

func (s *ApiSuggestIntegrationTest) TestCountry() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "росс",
	}
	res, err := api.Country(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestName() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "але",
	}
	res, err := api.Name(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestNameWithParts() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "але",
	}
	res, err := api.Name(context.Background(), &params, suggest.WithNameParts(suggest.NamePartSurname))
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestNameWithGender() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "але",
	}
	res, err := api.Name(context.Background(), &params, suggest.WithNameGender(suggest.NameGenderFemale))
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestEmail() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "ax",
	}
	res, err := api.Email(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiSuggestIntegrationTest) TestFMSUnit() {
	api := NewSuggestApi()
	params := suggest.RequestParams{
		Query: "увд",
	}
	res, err := api.FMSUnit(context.Background(), &params)
	s.NoError(err)
	s.NotEmpty(res)
}

func (s *ApiCleanIntegrationTest) TestAddress() {
	api := NewCleanApi()
	res, err := api.Address(context.Background(), "мск сухонска 11/-89")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestPhone() {
	api := NewCleanApi()
	res, err := api.Phone(context.Background(), "+79851234567")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestName() {
	api := NewCleanApi()
	res, err := api.Name(context.Background(), "Срегей владимерович иванов")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestEmail() {
	api := NewCleanApi()
	res, err := api.Email(context.Background(), "serega@yandex/ru")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestBirthday() {
	api := NewCleanApi()
	res, err := api.Birthday(context.Background(), "12 12 12")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestVehicle() {
	api := NewCleanApi()
	res, err := api.Vehicle(context.Background(), "форд фокус")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func (s *ApiCleanIntegrationTest) TestPassport() {
	api := NewCleanApi()
	res, err := api.Passport(context.Background(), "4509 235857")
	s.NoError(err)
	s.NotEmpty(res)
	s.Len(res, 1)
}

func ExampleNewSuggestApi() {
	api := NewSuggestApi()

	params := suggest.RequestParams{
		Query: "ул Свободы",
	}

	suggestions, err := api.Address(context.Background(), &params)
	if err != nil {
		return
	}

	for _, s := range suggestions {
		fmt.Printf("%s", s.Value)
	}
}
