package middleware

import (
	"errors"
	"net/http"
	"new-bloging-system/model"

	"github.com/gin-gonic/gin"
)

func PanicHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorResponse := model.Errors{
					Message: errors.New("unable to process the request. please try again after some time"),
					Type:    "internal_server_error",
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse)
				return
			}
		}()
		c.Next()
	}
}
