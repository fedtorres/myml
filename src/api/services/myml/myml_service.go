package myml

import (
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
)

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {
	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "userID is empty",
			Status: http.StatusBadRequest,
		}
	}
	user := myml.User {
		ID: int(userID),
	}
	if err := user.Get(); err != nil {
		return nil, &apierrors.ApiError {
		Message: "userID is empty",
		Status: http.StatusInternalServerError,
		}
	}
	return &user, nil
}
