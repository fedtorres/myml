package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

const urlUsers = "https://api.mercadolibre.com/users/"

func (user *User) Get() *apierrors.ApiError {
	final := fmt.Sprintf("%s%d", urlUsers, user.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status: http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status: http.StatusInternalServerError,
		}
	}
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return &apierrors.ApiError{
			Message: "userID is empty",
			Status: http.StatusInternalServerError,
		}
	}
	return nil
}
