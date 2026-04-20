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
	assert.Equal(t, FounderShareTypePercent, model.Founders[0].Share.Type)
}

func TestAddress(t *testing.T) {
	dataReader, err := os.Open("test/address.json")
	if err != nil {
		t.Error(err)
	}
	defer dataReader.Close()

	model := Address{}

	err = json.NewDecoder(dataReader).Decode(&model)
	assert.NoError(t, err)
	assert.Equal(t, "Москва, ул. Ленина, 1", model.Source)

	// sub_area fields (22.3+)
	assert.Equal(t, "sub-fias-id", model.SubAreaFiasID)
	assert.Equal(t, "sub-kladr-id", model.SubAreaKladrID)
	assert.Equal(t, "Одинцовское с/пос", model.SubAreaWithType)
	assert.Equal(t, "с/пос", model.SubAreaType)
	assert.Equal(t, "сельское поселение", model.SubAreaTypeFull)
	assert.Equal(t, "Одинцовское", model.SubArea)

	// stead fields (21.12+)
	assert.Equal(t, "stead-fias-id", model.SteadFiasID)
	assert.Equal(t, "50:20:0010101:123", model.SteadCadNum)
	assert.Equal(t, "уч", model.SteadType)
	assert.Equal(t, "участок", model.SteadTypeFull)
	assert.Equal(t, "5", model.Stead)

	// house_flat_count (24.3+)
	assert.Equal(t, "120", model.HouseFlatCount)

	// room fields (22.8+)
	assert.Equal(t, "room-fias-id", model.RoomFiasID)
	assert.Equal(t, "50:20:0010101:789", model.RoomCadNum)
	assert.Equal(t, "ком", model.RoomType)
	assert.Equal(t, "комната", model.RoomTypeFull)
	assert.Equal(t, "2", model.Room)

	// history_values
	assert.Equal(t, []string{"ул Сталина", "ул Советская"}, model.HistoryValues)

	// geoname_id
	assert.Equal(t, "524901", model.GeoNameId)
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
