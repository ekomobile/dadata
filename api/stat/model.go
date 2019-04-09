package stat

type (
	// Response usage statistics
	// see docs https://dadata.ru/api/stat/
	Stat struct {
		Date     string `json:"date"`
		Services *struct {
			Merging     int `json:"merging"`
			Suggestions int `json:"suggestions"`
			Clean       int `json:"clean"`
		} `json:"services"`
	}
)
