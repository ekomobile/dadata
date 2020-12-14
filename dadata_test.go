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
)

func (s *ApiSuggestIntegrationTest) SetupSuite() {
	if _, ok := os.LookupEnv("DADATA_API_KEY"); !ok {
		s.Suite.T().Skip("no api keys in env")
	}
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

func TestSuggestApiIntegration(t *testing.T) {
	suite.Run(t, &ApiSuggestIntegrationTest{})
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
