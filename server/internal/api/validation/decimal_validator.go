package validation

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(ValidateDecimal, decimal.Decimal{})
	}
}

func ValidateDecimal(field reflect.Value) any {
	if val, ok := field.Interface().(decimal.Decimal); ok {
		f, _ := val.Float64()
		return f
	}
	return nil
}
