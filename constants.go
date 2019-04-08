package dadata

// Определите, нужна ли дополнительная проверка оператором, используя код качества (qc):
const (
	QcSuccess = 0 // Исходное значение распознано уверенно. Не требуется ручная проверка.
	QcFailure = 1 // Исходное значение распознано с допущениями или не распознано. Требуется ручная проверка.
)

// Определите пригодность к рассылке, используя код полноты адреса (qc_complete):
const (
	QcCompleteSuitable               = 0  // Пригоден для почтовой рассылки
	QcCompleteNoRegion               = 1  // Не пригоден, нет региона
	QcCompleteNoCity                 = 2  // Не пригоден, нет города
	QcCompleteNoStreet               = 3  // Не пригоден, нет улицы
	QcCompleteNotHome                = 4  // Не пригоден, нет дома
	QcCompleteNoApartment            = 5  // Пригоден для юридических лиц или частных владений (нет квартиры)
	QcCompleteNotSuitable            = 6  // Не пригоден
	QcCompleteCompleteForeignAddress = 7  // Иностранный адрес
	QcCompleteCompleteNoKLADR        = 10 // Пригоден, но низкая вероятность успешной доставки (дом не найден в КЛАДР)
)

// Определите вероятность успешной доставки письма по адресу, используя код проверки дома (qc_house):
const (
	QcHouseExactMatch        = 2  // Дом найден по точному совпадению (КЛАДР)	Высокая
	QcHouseNotExpansionMatch = 3  // Различие в расширении дома (КЛАДР)	Средняя
	QcHouseRangeMatch        = 4  // Дом найден по диапазону (КЛАДР)	Средняя
	QcHouseNotFound          = 10 // Дом не найден (КЛАДР)	Низкая
)

// Определите точность координат адреса доставки с помощью кода qc_geo:
const (
	QcGeoExactCoordinates = 0 // Точные координаты
	QcGeoNearestHouse     = 1 // Ближайший дом
	QcGeoStreet           = 2 // Улица
	QcGeoLocality         = 3 // Населенный пункт
	QcGeoCity             = 4 // Город
	QcGeoNotDetermined    = 5 // Координаты не определены
)

// Проверьте, указал ли клиент телефон, соответствующий его адресу, с помощью кода qc_conflict (удобно для проверки уровня риска):
const (
	QcConflictFullMath   = 0 // Телефон соответствует адресу
	QcConflictCityMath   = 2 // Города адреса и телефона отличаются
	QcConflictRegionMath = 3 // Регионы адреса и телефона отличаются
)

// BoundValue type wrapper for suggest bounds
// full documentation https://confluence.hflabs.ru/pages/viewpage.action?pageId=222888017
type BoundValue string

// const for SuggestBound
const (
	SuggestBoundRegion     BoundValue = "region"     // Регион
	SuggestBoundArea       BoundValue = "area"       // Район
	SuggestBoundCity       BoundValue = "city"       // Город
	SuggestBoundSettlement BoundValue = "settlement" // Населенный пункт
	SuggestBoundStreet     BoundValue = "street"     // Улица
	SuggestBoundHouse      BoundValue = "house"      // Дом
)
