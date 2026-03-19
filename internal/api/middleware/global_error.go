package middleware

import (
	"errors"
	"io"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/pkg/response"
	"github.com/BramAristyo/go-pos-mawish/pkg/usecase_errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if errors.Is(err, io.EOF) {
				response.Error(c, http.StatusBadRequest, "request body is empty", err)
				return
			}

			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				response.ValidationError(c, err)
				return
			}

			var ue *usecase_errors.UseCaseError
			if errors.As(err, &ue) {
				response.Error(c, ue.Code, ue.Message, err)
				return
			}

			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(c, http.StatusNotFound, "resource not found", err)
				return
			}

			if usecase_errors.IsUniqueViolation(err) {
				response.Error(c, http.StatusConflict, "data already exists", err)
				return
			}

			response.Error(c, http.StatusInternalServerError, err.Error(), err)
		}
	}
}
