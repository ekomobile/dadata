package model

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank_Recursive(t *testing.T) {
	testBankReader, err := os.Open("test/bank.json")
	if err != nil {
		t.Error(err)
	}
	defer testBankReader.Close()

	bank := Bank{}

	err = json.NewDecoder(testBankReader).Decode(&bank)
	assert.NoError(t, err)
	assert.Equal(t, "111", bank.Bic)
	assert.Equal(t, "222", bank.Rkc.Bic)
	assert.Equal(t, "333", bank.Rkc.Rkc.Bic)
	assert.Equal(t, "444", bank.Rkc.Rkc.Rkc.Bic)
}

func TestParty(t *testing.T) {
	testPartyReader, err := os.Open("test/party.json")
	if err != nil {
		t.Error(err)
	}
	defer testPartyReader.Close()

	model := Party{}

	err = json.NewDecoder(testPartyReader).Decode(&model)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(model.Founders))
	assert.Equal(t, PartyFounderTypeLegal, model.Founders[0].Type)
	assert.Equal(t, PartySMBCategorySmall, model.Documents.Smb.Category)
}

func TestPhone1(t *testing.T) {
	dataReader, err := os.Open("test/phone1.json")
	if err != nil {
		t.Error(err)
	}
	defer dataReader.Close()

	model := Phone{}

	err = json.NewDecoder(dataReader).Decode(&model)
	assert.NoError(t, err)
	assert.Equal(t, "+375123456789", model.Source)
	assert.Equal(t, "Беларусь", model.Country)
	assert.Equal(t, "", model.City)
}

func TestPhone2(t *testing.T) {
	dataReader, err := os.Open("test/phone2.json")
	if err != nil {
		t.Error(err)
	}
	defer dataReader.Close()

	model := Phone{}

	err = json.NewDecoder(dataReader).Decode(&model)
	assert.NoError(t, err)
	assert.Equal(t, "+79851234567", model.Source)
	assert.Equal(t, "Россия", model.Country)
	assert.Equal(t, "", model.City)
}

func TestPhone3(t *testing.T) {
	dataReader, err := os.Open("test/phone3.json")
	if err != nil {
		t.Error(err)
	}
	defer dataReader.Close()

	model := Phone{}

	err = json.NewDecoder(dataReader).Decode(&model)
	assert.NoError(t, err)
	assert.Equal(t, "+123", model.Source)
	assert.Equal(t, "", model.Country)
	assert.Equal(t, "", model.City)
}
