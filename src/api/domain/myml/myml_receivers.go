package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
)

const urlUsers = "http://localhost:8081/users/"

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

func (user *User) Get() *apierrors.ApiError {
	final := fmt.Sprintf("%s%d", urlUsers, user.ID)
	data, err := cb.Execute(func() (interface{}, error) {
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
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	if err := json.Unmarshal([]byte(data.([]byte)), &user); err != nil {
		return &apierrors.ApiError{
			Message: "json.Unmarshal failed",
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}
