package api

import (
	"github.com/NovikovAndrew/SimpleBank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
