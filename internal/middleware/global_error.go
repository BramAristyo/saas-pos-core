package middleware

import (
	"errors"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				response.ValidationError(c, err)
				return
			}

			var se *service_errors.ServiceError
			if errors.As(err, &se) {
				response.Error(c, se.Code, se.Message)
				return
			}

			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(c, http.StatusNotFound, "resource not found")
				return
			}

			if service_errors.IsUniqueViolation(err) {
				response.Error(c, http.StatusConflict, "data already exists")
				return
			}

			response.Error(c, http.StatusInternalServerError, err.Error())
		}
	}
}
