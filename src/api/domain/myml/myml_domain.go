package myml

import "github.com/mercadolibre/myml/src/api/utils/apierrors"

type User struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Address          struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	Logo             interface{} `json:"logo"`
	Points           int         `json:"points"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}

type Categories []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type Currency struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	Symbol        string `json:"symbol"`
	DecimalPlaces int    `json:"decimal_places"`
}

type MyML struct {
	Categories Categories          `json:"categories"`
	Currency   Currency            `json:"currency"`
	Error      *apierrors.ApiError `json:"error"`
}
