package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
)

const (
	paramUserID = "userID"
)

func GetUser(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
