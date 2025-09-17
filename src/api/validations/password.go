package validations

import (
	"github.com/rezamobaraki/CarApp/common"
	"github.com/go-playground/validator/v10"
)

func ValidatePassword(validationField validator.FieldLevel) bool {
	value, isValidType := validationField.Field().Interface().(string)
	if !isValidType {
		return false
	}
	return common.CheckPassword(value)
}
