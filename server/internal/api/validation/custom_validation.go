package validation

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		var el ValidationError
		el.Property = err.Field()
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
		ve.Message = ve.Property + " is required"
	case "min":
		ve.Message = ve.Property + " must be at least " + ve.Value
	case "max":
		ve.Message = ve.Property + " must be at most " + ve.Value
	case "email":
		ve.Message = ve.Property + " must be a valid email address"

	case "gte":
		ve.Message = ve.Property + " must be greater than or equal to " + ve.Value
	case "lte":
		ve.Message = ve.Property + " must be less than or equal to " + ve.Value
	case "numeric":
		ve.Message = ve.Property + " must be a valid number"
	case "oneof":
		ve.Message = ve.Property + " must be one of: [" + ve.Value + "]"
	case "len":
		ve.Message = ve.Property + " must be exactly " + ve.Value + " characters/items"
	case "uuid":
		ve.Message = ve.Property + " must be a valid UUID format"

	default:
		ve.Message = ve.Property + " is invalid"
	}
}
