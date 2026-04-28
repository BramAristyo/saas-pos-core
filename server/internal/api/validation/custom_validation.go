package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property        string `json:"property"`
	Tag             string `json:"tag"`
	Value           string `json:"value"`
	Message         string `json:"message"`
	PropertyMessage string `json:"-"`
}

func GetValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var el ValidationError

		namespace := err.StructNamespace()

		parts := strings.SplitN(namespace, ".", 2)
		if len(parts) == 2 {
			el.Property = parts[1]
		} else {
			el.Property = namespace
		}

		el.PropertyMessage = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		el.setValidationMessage()
		validationErrors = append(validationErrors, el)
	}

	return validationErrors
}

func (ve *ValidationError) setValidationMessage() {

	switch ve.Tag {
	case "required":
		ve.Message = ve.PropertyMessage + " is required"
	case "min":
		ve.Message = ve.PropertyMessage + " must be at least " + ve.Value
	case "max":
		ve.Message = ve.PropertyMessage + " must be at most " + ve.Value
	case "email":
		ve.Message = ve.PropertyMessage + " must be a valid email address"

	case "gt":
		ve.Message = ve.PropertyMessage + " must be greater than " + ve.Value
	case "lt":
		ve.Message = ve.PropertyMessage + " must be less than " + ve.Value
	case "gte":
		ve.Message = ve.PropertyMessage + " must be greater than or equal to " + ve.Value
	case "lte":
		ve.Message = ve.PropertyMessage + " must be less than or equal to " + ve.Value
	case "numeric":
		ve.Message = ve.PropertyMessage + " must be a valid number"
	case "oneof":
		ve.Message = ve.PropertyMessage + " must be one of: [" + ve.Value + "]"
	case "len":
		ve.Message = ve.PropertyMessage + " must be exactly " + ve.Value + " characters/items"
	case "uuid":
		ve.Message = ve.PropertyMessage + " must be a valid UUID format"

	default:
		ve.Message = ve.PropertyMessage + " is invalid"
	}
}
