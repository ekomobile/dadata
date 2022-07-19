package model

// Определите, нужна ли дополнительная проверка оператором, используя код качества (qc):
const (
	QCSuccess = 0 // Исходное значение распознано уверенно. Не требуется ручная проверка.
	QCFailure = 1 // Исходное значение распознано с допущениями или не распознано. Требуется ручная проверка.
)

// Определите пригодность к рассылке, используя код полноты адреса (qc_complete):
const (
	QCCompleteSuitable               = 0  // Пригоден для почтовой рассылки
	QCCompleteNoRegion               = 1  // Не пригоден, нет региона
	QCCompleteNoCity                 = 2  // Не пригоден, нет города
	QCCompleteNoStreet               = 3  // Не пригоден, нет улицы
	QCCompleteNotHome                = 4  // Не пригоден, нет дома
	QCCompleteNoApartment            = 5  // Пригоден для юридических лиц или частных владений (нет квартиры)
	QCCompleteNotSuitable            = 6  // Не пригоден
	QCCompleteCompleteForeignAddress = 7  // Иностранный адрес
	QCCompleteCompleteNoKLADR        = 10 // Пригоден, но низкая вероятность успешной доставки (дом не найден в КЛАДР)
)

// Определите вероятность успешной доставки письма по адресу, используя код проверки дома (qc_house):
const (
	QCHouseExactMatch        = 2  // Дом найден по точному совпадению (КЛАДР)	Высокая
	QCHouseNotExpansionMatch = 3  // Различие в расширении дома (КЛАДР)	Средняя
	QCHouseRangeMatch        = 4  // Дом найден по диапазону (КЛАДР)	Средняя
	QCHouseNotFound          = 10 // Дом не найден (КЛАДР)	Низкая
)

// Определите точность координат адреса доставки с помощью кода qc_geo:
const (
	QCGeoExactCoordinates = iota // Точные координаты
	QCGeoNearestHouse            // Ближайший дом
	QCGeoStreet                  // Улица
	QCGeoLocality                // Населенный пункт
	QCGeoCity                    // Город
	QCGeoNotDetermined           // Координаты не определены
)

// Проверьте, указал ли клиент телефон, соответствующий его адресу, с помощью кода qc_conflict (удобно для проверки уровня риска):
const (
	QCConflictFullMath   = 0 // Телефон соответствует адресу
	QCConflictCityMath   = 2 // Города адреса и телефона отличаются
	QCConflictRegionMath = 3 // Регионы адреса и телефона отличаются
)

// const for SuggestBound
const (
	SuggestBoundRegion     BoundValue = "region"     // Регион
	SuggestBoundArea       BoundValue = "area"       // Район
	SuggestBoundCity       BoundValue = "city"       // Город
	SuggestBoundSettlement BoundValue = "settlement" // Населенный пункт
	SuggestBoundStreet     BoundValue = "street"     // Улица
	SuggestBoundHouse      BoundValue = "house"      // Дом
)

// FMS unit type
// https://dadata.ru/api/suggest/fms_unit/
const (
	FMSTypeFMS         = 0 // Подразделение ФМС
	FMSTypeMVD         = 1 // ГУВД или МВД региона
	FMSTypeOVD         = 2 // УВД или ОВД района или города
	FMSTypePoliceState = 3 // Отделение полиции
)

type (
	// BoundValue type wrapper for suggest bounds
	// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=222888017
	BoundValue string

	// Address base struct for datdata.Address
	Address struct {
		Source               string `json:"source"`                  // Исходный адрес одной строкой
		Result               string `json:"result"`                  // Стандартизованный адрес одной строкой
		PostalCode           string `json:"postal_code"`             // Индекс
		Country              string `json:"country"`                 // Страна
		CountryIsoCode       string `json:"country_iso_code"`        // ISO-код страны
		FederalDistrict      string `json:"federal_district"`        // Федеральный округ
		RegionFiasID         string `json:"region_fias_id"`          // Код ФИАС региона
		RegionKladrID        string `json:"region_kladr_id"`         // Код КЛАДР региона
		RegionIsoCode        string `json:"region_iso_code"`         // ISO-код региона
		RegionWithType       string `json:"region_with_type"`        // Регион с типом
		RegionType           string `json:"region_type"`             // Тип региона (сокращенный)
		RegionTypeFull       string `json:"region_type_full"`        // Тип региона
		Region               string `json:"region"`                  // Регион
		AreaFiasID           string `json:"area_fias_id"`            // Код ФИАС района в регионе
		AreaKladrID          string `json:"area_kladr_id"`           // Код КЛАДР района в регионе
		AreaWithType         string `json:"area_with_type"`          // Район в регионе с типом
		AreaType             string `json:"area_type"`               // Тип района в регионе (сокращенный)
		AreaTypeFull         string `json:"area_type_full"`          // Тип района в регионе
		Area                 string `json:"area"`                    // Район в регионе
		CityFiasID           string `json:"city_fias_id"`            // Код ФИАС города
		CityKladrID          string `json:"city_kladr_id"`           // Код КЛАДР города
		CityWithType         string `json:"city_with_type"`          // Город с типом
		CityType             string `json:"city_type"`               // Тип города (сокращенный)
		CityTypeFull         string `json:"city_type_full"`          // Тип города
		City                 string `json:"city"`                    // Город
		CityArea             string `json:"city_area"`               // Административный округ (только для Москвы)
		CityDistrictFiasID   string `json:"city_district_fias_id"`   // Код ФИАС района города (заполняется, только если район есть в ФИАС)
		CityDistrictKladrID  string `json:"city_district_kladr_id"`  // Код КЛАДР района города (не заполняется)
		CityDistrictWithType string `json:"city_district_with_type"` // Район города с типом
		CityDistrictType     string `json:"city_district_type"`      // Тип района города (сокращенный)
		CityDistrictTypeFull string `json:"city_district_type_full"` // Тип района города
		CityDistrict         string `json:"city_district"`           // Район города
		SettlementFiasID     string `json:"settlement_fias_id"`      // Код ФИАС нас. пункта
		SettlementKladrID    string `json:"settlement_kladr_id"`     // Код КЛАДР нас. пункта
		SettlementWithType   string `json:"settlement_with_type"`    // Населенный пункт с типом
		SettlementType       string `json:"settlement_type"`         // Тип населенного пункта (сокращенный)
		SettlementTypeFull   string `json:"settlement_type_full"`    // Тип населенного пункта
		Settlement           string `json:"settlement"`              // Населенный пункт
		StreetFiasID         string `json:"street_fias_id"`          // Код ФИАС улицы
		StreetKladrID        string `json:"street_kladr_id"`         // Код КЛАДР улицы
		StreetWithType       string `json:"street_with_type"`        // Улица с типом
		StreetType           string `json:"street_type"`             // Тип улицы (сокращенный)
		StreetTypeFull       string `json:"street_type_full"`        // Тип улицы
		Street               string `json:"street"`                  // Улица
		HouseFiasID          string `json:"house_fias_id"`           // Код ФИАС дома
		HouseKladrID         string `json:"house_kladr_id"`          // Код КЛАДР дома
		HouseType            string `json:"house_type"`              // Тип дома (сокращенный)
		HouseTypeFull        string `json:"house_type_full"`         // Тип дома
		House                string `json:"house"`                   // Дом
		BlockType            string `json:"block_type"`              // Тип корпуса/строения (сокращенный)
		BlockTypeFull        string `json:"block_type_full"`         // Тип корпуса/строения
		Block                string `json:"block"`                   // Корпус/строение
		Entrance             string `json:"entrance"`                // Подъезд
		Floor                string `json:"floor"`                   // Этаж
		FlatFiasId           string `json:"flat_fias_id"`            // ФИАС-код квартиры
		FlatType             string `json:"flat_type"`               // Тип квартиры (сокращенный)
		FlatTypeFull         string `json:"flat_type_full"`          // Тип квартиры
		Flat                 string `json:"flat"`                    // Квартира
		FlatArea             string `json:"flat_area"`               // Площадь квартиры
		SquareMeterPrice     string `json:"square_meter_price"`      // Рыночная стоимость м²
		FlatPrice            string `json:"flat_price"`              // Рыночная стоимость квартиры
		PostalBox            string `json:"postal_box"`              // Абонентский ящик
		FiasID               string `json:"fias_id"`                 // Код ФИАС
		FiasCode             string `json:"fias_code"`               // Иерархический код адреса в ФИАС (СС+РРР+ГГГ+ППП+СССС+УУУУ+ДДДД)
		FiasLevel            string `json:"fias_level"`              // Уровень детализации, до которого адрес найден в ФИАС
		FiasActualityState   string `json:"fias_actuality_state"`    // Признак актуальности адреса в ФИАС
		KladrID              string `json:"kladr_id"`                // Код КЛАДР
		CapitalMarker        string `json:"capital_marker"`          // Статус центра
		Okato                string `json:"okato"`                   // Код ОКАТО
		Oktmo                string `json:"oktmo"`                   // Код ОКТМО
		TaxOffice            string `json:"tax_office"`              // Код ИФНС для физических лиц
		TaxOfficeLegal       string `json:"tax_office_legal"`        // Код ИФНС для организаций
		Timezone             string `json:"timezone"`                // Часовой пояс
		GeoLat               string `json:"geo_lat"`                 // Координаты: широта
		GeoLon               string `json:"geo_lon"`                 // Координаты: долгота
		BeltwayHit           string `json:"beltway_hit"`             // Внутри кольцевой?
		BeltwayDistance      string `json:"beltway_distance"`        // Расстояние от кольцевой в км.
		// QualityCodeGeoRaw для clean вызовов он int для suggest в адресе банков он string поэтому в поле поставил interface{} чтобы работало и там и там)\
		QualityCodeGeoRaw      interface{} `json:"qc_geo"`         // Код точности координат
		QualityCodeCompleteRaw interface{} `json:"qc_complete"`    // Код полноты
		QualityCodeHouseRaw    interface{} `json:"qc_house"`       // Код проверки дома
		QualityCodeRaw         interface{} `json:"qc"`             // Код качества
		UnparsedParts          string      `json:"unparsed_parts"` // Нераспознанная часть адреса. Для адреса
		Metro                  []*Metro    `json:"metro"`
	}

	// AddressResponse api response for address
	AddressResponse struct {
		Value             string   `json:"value"`
		UnrestrictedValue string   `json:"unrestricted_value"`
		Data              *Address `json:"data"`
	}

	// Metro base struct for dadata.Metro
	Metro struct {
		Name     string  `json:"name"`
		Line     string  `json:"line"`
		Distance float64 `json:"distance"`
	}

	// Phone base struct for dadata.Phone
	Phone struct {
		Source              string `json:"source"`       // Исходный телефон одной строкой
		Type                string `json:"type"`         // Тип телефона
		Phone               string `json:"phone"`        // Стандартизованный телефон одной строкой
		CountryCode         string `json:"country_code"` // Код страны
		CityCode            string `json:"city_code"`    // Код города / DEF-код
		Number              string `json:"number"`       // Локальный номер телефона
		Extension           string `json:"extension"`    // Добавочный номер
		Provider            string `json:"provider"`     // Оператор связи
		Region              string `json:"region"`       // Регион
		Timezone            string `json:"timezone"`     // Часовой пояс
		QualityCodeConflict int    `json:"qc_conflict"`  // Признак конфликта телефона с адресом
		QualityCode         int    `json:"qc"`           // Код качества
	}

	// Name base struct for dadata.Name
	Name struct {
		Source         string      `json:"source"`          // Исходное ФИО одной строкой
		Result         string      `json:"result"`          // Стандартизованное ФИО одной строкой
		ResultGenitive string      `json:"result_genitive"` // ФИО в родительном падеже (кого?)
		ResultDative   string      `json:"result_dative"`   // ФИО в дательном падеже (кому?)
		ResultAblative string      `json:"result_ablative"` // ФИО в творительном падеже (кем?)
		Surname        string      `json:"surname"`         // Фамилия
		Name           string      `json:"name"`            // Имя
		Patronymic     string      `json:"patronymic"`      // Отчество
		Gender         string      `json:"gender"`          // Пол
		QualityCodeRaw interface{} `json:"qc"`              // Код качества
	}

	// Email base struct for dadata.Email
	Email struct {
		Source      string `json:"source"` // Исходный e-mail
		Email       string `json:"email"`  // Стандартизованный e-mail
		QualityCode int    `json:"qc"`     // Код качества
	}

	// Birthday base struct for dadata.Birthday
	Birthday struct {
		Source      string `json:"source"`    // Исходная дата
		Birthdate   string `json:"birthdate"` // Стандартизованная дата
		QualityCode int    `json:"qc"`        // Код качества
	}

	// Vehicle base struct for dadata.Vehicle
	Vehicle struct {
		Source      string `json:"source"` // Исходное значение
		Result      string `json:"result"` // Стандартизованное значение
		Brand       string `json:"brand"`  // Марка
		Model       string `json:"model"`  // Модель
		QualityCode int    `json:"qc"`     // Код проверки
	}

	// Passport base struct for dadata.Passport
	Passport struct {
		Source      string `json:"source"` // Исходная серия и номер одной строкой
		Series      string `json:"series"` // Серия
		Number      string `json:"number"` // Номер
		QualityCode int    `json:"qc"`     // Код проверки
	}

	// Bank base struct for dadata.Bank
	Bank struct {
		Opf                  *OrganizationOPF   `json:"opf"`
		Name                 *BankName          `json:"name"`
		Bic                  string             `json:"bic"`                   // Банковский идентификационный код (БИК) ЦБ РФ
		Swift                string             `json:"swift"`                 // Банковский идентификационный код в системе SWIFT
		Okpo                 string             `json:"okpo"`                  // Код ОКПО
		CorrespondentAccount string             `json:"correspondent_account"` // Корреспондентский счет в ЦБ РФ
		RegistrationNumber   string             `json:"registration_number"`   // Регистрационный номер в ЦБ РФ
		Rkc                  *Bank              `json:"rkc"`                   // Расчетно-кассовый центр. Объект такой же структуры, как сам банк.
		Address              *AddressResponse   `json:"address"`               // см AddressResponse
		Phone                string             `json:"phone"`                 // Не заполняется
		State                *OrganizationState `json:"state"`
	}

	// OrganizationOPF Тип Кредитной организации
	OrganizationOPF struct {
		Code  string `json:"code"`  // код ОКОПФ
		Type  string `json:"type"`  // Тип кредитной организации
		Full  string `json:"full"`  // Тип кредитной организации (на русском)
		Short string `json:"short"` // Тип кредитной организации (на русском, сокращенный)
	}

	// BankName наименование банка
	BankName struct {
		Payment string `json:"payment"` // Платежное наименование
		Full    string `json:"full"`    // Полное наименование
		Short   string `json:"short"`   // Краткое наименование
	}

	// OrganizationState Статус организации
	OrganizationState struct {
		Status string `json:"status"` // Статус организации:
		//  ACTIVE      — действующая
		//  LIQUIDATING — ликвидируется
		//  LIQUIDATED  — ликвидирована
		ActualityDate    int64 `json:"actuality_date"`    // Дата актуальности сведений
		RegistrationDate int64 `json:"registration_date"` // Дата регистрации
		LiquidationDate  int64 `json:"liquidation_date"`  // Дата ликвидации
	}

	// Party base struct for dadata.Party (rus Организация)
	Party struct {
		Kpp        string `json:"kpp"`
		Capital    string `json:"capital"`
		Management *struct {
			Name string `json:"name"`
			Post string `json:"post"`
		} `json:"management"`
		Founders    string             `json:"founders"`
		Managers    string             `json:"managers"`
		BranchType  string             `json:"branch_type"`
		BranchCount int                `json:"branch_count"`
		Source      string             `json:"source"`
		QC          string             `json:"qc"`
		Hid         string             `json:"hid"`
		Type        string             `json:"type"`
		State       *OrganizationState `json:"state"`
		Opf         *OrganizationOPF   `json:"opf"`
		Name        *struct {
			FullWithOpf  string `json:"full_with_opf"`
			ShortWithOpf string `json:"short_with_opf"`
			Latin        string `json:"latin"`
			Full         string `json:"full"`
			Short        string `json:"short"`
		} `json:"name"`
		Inn    string `json:"inn"`
		Ogrn   string `json:"ogrn"`
		Okato  string `json:"okato"`
		Oktmo  string `json:"oktmo"`
		Okpo   string `json:"okpo"`
		Okogu  string `json:"okogu"`
		Okfs   string `json:"okfs"`
		Okved  string `json:"okved"`
		Okveds []*struct {
			Main bool   `json:"main"`
			Type string `json:"type"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"okveds"`
		Authorities string           `json:"authorities"`
		Documents   string           `json:"documents"`
		Licenses    string           `json:"licenses"`
		Address     *AddressResponse `json:"address"`
		Phones      string           `json:"phones"`
		Emails      string           `json:"emails"`
		OgrnDate    int64            `json:"ogrn_date"`
		OkvedType   string           `json:"okved_type"`
	}

	// Country base struct for dadata.Country
	Country struct {
		Code      string `json:"code"`
		Alfa2     string `json:"alfa2"`
		Alfa3     string `json:"alfa3"`
		NameShort string `json:"name_short"`
		Name      string `json:"name"`
	}

	// FMSUnit is a FMS unit data model
	// https://dadata.ru/api/suggest/fms_unit/
	FMSUnit struct {
		Code       string `json:"code"`
		Name       string `json:"name"`
		RegionCode string `json:"region_code"`
		Type       string `json:"type"`
	}
)
