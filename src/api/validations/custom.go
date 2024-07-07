package validations

import (
	"github.com/MrRezoo/CarApp/common"
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

func ValidateIranianMobile(validationField validator.FieldLevel) bool {
	mobileNumber, isValidType := validationField.Field().Interface().(string)
	if !isValidType {
		return false
	}
	isValidFormat, err := regexp.MatchString(`^(\+98|0)?9\d{9}$`, mobileNumber)
	if err != nil {
		log.Println("Error validating Iranian mobile number: ", err)
	}
	return isValidFormat
}

func ValidatePassword(validationField validator.FieldLevel) bool {
	value, isValidType := validationField.Field().Interface().(string)
	if !isValidType {
		return false
	}
	return common.CheckPassword(value)
}
