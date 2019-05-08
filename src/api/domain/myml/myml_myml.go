package myml

import "github.com/mercadolibre/myml/src/api/utils/apierrors"

type MyML struct {
	Categories *Categories         `json:"categories"`
	Currency   *Currency           `json:"currency"`
	Error      *apierrors.ApiError `json:"error"`
}
