package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

const urlCategories = "https://api.mercadolibre.com/sites/"
const urlCountries = "https://api.mercadolibre.com/classified_locations/countries/"
const urlCurrencies = "https://api.mercadolibre.com/currencies/"

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {
	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "userID is empty",
			Status:  http.StatusBadRequest,
		}
	}
	user := myml.User{
		ID: int(userID),
	}
	if err := user.Get(); err != nil {
		return nil, &apierrors.ApiError{
			Message: "user.Get failed",
			Status:  http.StatusInternalServerError,
		}
	}
	return &user, nil
}

func GetMyMLFromAPI(userID int64) (*myml.MyML, *apierrors.ApiError) {
	user, err := GetUserFromAPI(userID)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: "GetUserFromAPI failed",
			Status:  http.StatusInternalServerError,
		}
	}
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan *myml.MyML)
	go func() { c <- getCategories(user.SiteID, &wg) }()
	go func() { c <- getCurrency(user.CountryID, &wg) }()
	wg.Wait()
	var myML myml.MyML
	for elem := range c {
		if elem.Categories != nil {
			myML.Categories = elem.Categories
			continue
		}
		if &elem.Currency != nil {
			myML.Currency = elem.Currency
		}
		if elem.Error != nil {
			return nil, &apierrors.ApiError{
				Message: "Couldn't get reply",
				Status:  http.StatusInternalServerError,
			}
		}
	}
	return &myML, nil
}

func getCategories(countryID string, wg *sync.WaitGroup) *myml.MyML {
	defer wg.Done()
	final := fmt.Sprintf("%s%s/categories", urlCategories, countryID)
	fmt.Println(final)
	response, err := http.Get(final)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	var categories myml.Categories
	if err := json.Unmarshal([]byte(data), &categories); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	myML := myml.MyML{
		Categories: categories,
	}
	fmt.Println(myML)
	return &myML
}

func getCurrency(countryID string, wg *sync.WaitGroup) *myml.MyML {
	defer wg.Done()
	final := fmt.Sprintf("%s%s", urlCountries, countryID)
	fmt.Println(final)
	response, err := http.Get(final)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	country := myml.Country{}
	if err := json.Unmarshal([]byte(data), &country); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	final = fmt.Sprintf("%s%s", urlCurrencies, country.CurrencyID)
	fmt.Println(final)
	response, err = http.Get(final)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	var currency myml.Currency
	if err := json.Unmarshal([]byte(data), &currency); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "http.Get failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	myML := myml.MyML{
		Currency: currency,
	}
	fmt.Println(myML.Currency)
	return &myML
}
