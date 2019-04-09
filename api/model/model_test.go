package model

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank_Recursive(t *testing.T) {
	srcData := `{
		"opf":                   {
			"type":  "BANK",
			"full":  null,
			"short": null
		},
		"name":                  {
			"payment": "ПАО СБЕРБАНК",
			"full":    null,
			"short":   "СБЕРБАНК РОССИИ"
		},
		"bic":                   "111",
		"swift":                 "SABRRUMMXXX",
		"okpo":                  null,
		"correspondent_account": "30101810400000000225",
		"registration_number":   "1481",
		"rkc":                   {
			"opf":                   {
				"type":  "BANK",
				"full":  null,
				"short": null
			},
			"name":                  {
				"payment": "ПАО СБЕРБАНК",
				"full":    null,
				"short":   "СБЕРБАНК РОССИИ"
			},
			"bic":                   "222",
			"swift":                 "SABRRUMMXXX",
			"okpo":                  null,
			"correspondent_account": "30101810400000000225",
			"registration_number":   "1481",
			"rkc":                   {
				"opf":                   {
					"type":  "BANK",
					"full":  null,
					"short": null
				},
				"name":                  {
					"payment": "ПАО СБЕРБАНК",
					"full":    null,
					"short":   "СБЕРБАНК РОССИИ"
				},
				"bic":                   "333",
				"swift":                 "SABRRUMMXXX",
				"okpo":                  null,
				"correspondent_account": "30101810400000000225",
				"registration_number":   "1481",
				"rkc":                   {
					"opf":                   {
						"type":  "BANK",
						"full":  null,
						"short": null
					},
					"name":                  {
						"payment": "ПАО СБЕРБАНК",
						"full":    null,
						"short":   "СБЕРБАНК РОССИИ"
					},
					"bic":                   "444",
					"swift":                 "SABRRUMMXXX",
					"okpo":                  null,
					"correspondent_account": "30101810400000000225",
					"registration_number":   "1481",
					"rkc":                   null,
					"address":               {
						"value":              "г Москва, ул Вавилова, д 19",
						"unrestricted_value": "г Москва, Академический р-н, ул Вавилова, д 19",
						"data":               {}
					},
					"phones":                null,
					"state":                 {
						"status":            "ACTIVE",
						"actuality_date":    1548201600000,
						"registration_date": 677376000000,
						"liquidation_date":  null
					}
				},
				"address":               {
					"value":              "г Москва, ул Вавилова, д 19",
					"unrestricted_value": "г Москва, Академический р-н, ул Вавилова, д 19",
					"data":               {}
				},
				"phones":                null,
				"state":                 {
					"status":            "ACTIVE",
					"actuality_date":    1548201600000,
					"registration_date": 677376000000,
					"liquidation_date":  null
				}
			},
			"address":               {
				"value":              "г Москва, ул Вавилова, д 19",
				"unrestricted_value": "г Москва, Академический р-н, ул Вавилова, д 19",
				"data":               {}
			},
			"phones":                null,
			"state":                 {
				"status":            "ACTIVE",
				"actuality_date":    1548201600000,
				"registration_date": 677376000000,
				"liquidation_date":  null
			}
		},
		"address":               {
			"value":              "г Москва, ул Вавилова, д 19",
			"unrestricted_value": "г Москва, Академический р-н, ул Вавилова, д 19",
			"data":               {}
		},
		"phones":                null,
		"state":                 {
			"status":            "ACTIVE",
			"actuality_date":    1548201600000,
			"registration_date": 677376000000,
			"liquidation_date":  null
		}
	}`

	bank := Bank{}

	err := json.NewDecoder(strings.NewReader(srcData)).Decode(&bank)
	assert.NoError(t, err)
	assert.Equal(t, "111", bank.Bic)
	assert.Equal(t, "222", bank.Rkc.Bic)
	assert.Equal(t, "333", bank.Rkc.Rkc.Bic)
	assert.Equal(t, "444", bank.Rkc.Rkc.Rkc.Bic)
}
