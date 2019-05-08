package myml

type Country struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Locale             string      `json:"locale"`
	CurrencyID         string      `json:"currency_id"`
	DecimalSeparator   string      `json:"decimal_separator"`
	ThousandsSeparator string      `json:"thousands_separator"`
	TimeZone           string      `json:"time_zone"`
	GeoInformation     interface{} `json:"geo_information"`
	States             []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}
