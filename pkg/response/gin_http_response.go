package response

import (
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BaseHTTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func OK(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, BaseHTTPResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func Created(c *gin.Context, data any, message string) {
	c.JSON(http.StatusCreated, BaseHTTPResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, BaseHTTPResponse{
		Success: false,
		Error:   message,
	})
}

func ValidationError(c *gin.Context, err validator.ValidationErrors) {
	c.JSON(http.StatusBadRequest, BaseHTTPResponse{
		Success: false,
		Error:   validation.GetValidationErrors(err),
	})
}
