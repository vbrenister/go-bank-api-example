package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/vbrenister/go-bank-api-example/util"
)

var validatorCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupporterCurrency(currency)
	}

	return false

}
