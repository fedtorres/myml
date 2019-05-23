package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
	"sync"
)

const mockUrl = "http://localhost:8081"
const urlCategories = mockUrl + "/sites/"
const urlCountries = mockUrl + "/classified_locations/countries/"
const urlCurrencies = mockUrl + "/currencies/"

var cbCategories *gobreaker.CircuitBreaker
var cbCurrency *gobreaker.CircuitBreaker
var cbCountry *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cbCategories = gobreaker.NewCircuitBreaker(st)
	cbCurrency = gobreaker.NewCircuitBreaker(st)
	cbCountry = gobreaker.NewCircuitBreaker(st)
}

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
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return &user, nil
}

func GetMyMLFromAPI(userID int64) (*myml.MyML, *apierrors.ApiError) {
	user, err := GetUserFromAPI(userID)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan *myml.MyML, 2)
	go func() { c <- getCategories(user.SiteID, &wg) }()
	go func() { c <- getCurrency(user.CountryID, &wg) }()
	wg.Wait()
	close(c)
	var myML myml.MyML
	for elem := range c {
		if elem.Categories != nil {
			myML.Categories = elem.Categories
			continue
		}
		if elem.Currency != nil {
			myML.Currency = elem.Currency
			continue
		}
		if elem.Error != nil {
			return &myML, &apierrors.ApiError{
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
	data, err := cbCategories.Execute(func() (interface{}, error) {
		response, err := http.Get(final)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return data, nil
	})
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}
	var categories myml.Categories
	if err := json.Unmarshal([]byte(data.([]byte)), &categories); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "json.Unmarshal failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	myML := myml.MyML{
		Categories: &categories,
	}
	fmt.Println(myML)
	return &myML
}

func getCurrency(countryID string, wg *sync.WaitGroup) *myml.MyML {
	defer wg.Done()
	final := fmt.Sprintf("%s%s", urlCountries, countryID)
	data, err := cbCountry.Execute(func() (interface{}, error) {
		response, err := http.Get(final)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return data, nil
	})
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}
	country := myml.Country{}
	if err := json.Unmarshal([]byte(data.([]byte)), &country); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "json.Unmarshal failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	final = fmt.Sprintf("%s%s", urlCurrencies, country.CurrencyID)
	data, err = cbCurrency.Execute(func() (interface{}, error) {
		response, err := http.Get(final)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return data, nil
	})
	if err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}
	var currency myml.Currency
	if err := json.Unmarshal([]byte(data.([]byte)), &currency); err != nil {
		return &myml.MyML{
			Categories: nil,
			Currency:   nil,
			Error: &apierrors.ApiError{
				Message: "json.Unmarshal failed.",
				Status:  http.StatusInternalServerError,
			},
		}
	}
	myML := myml.MyML{
		Currency: &currency,
	}
	fmt.Println(myML)
	return &myML
}
