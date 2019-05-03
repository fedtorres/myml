package myml

type User struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Address          struct {
	} `json:"address"`
	UserType         string        `json:"user_type"`
	Tags             []interface{} `json:"tags"`
	Logo             interface{}   `json:"logo"`
	Points           int           `json:"points"`
	SiteID           string        `json:"site_id"`
	Permalink        string        `json:"permalink"`
	SellerReputation struct {
	} `json:"seller_reputation"`
	BuyerReputation struct {
	} `json:"buyer_reputation"`
	Status struct {
	} `json:"status"`
}