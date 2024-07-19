package helper

import (
	"github.com/MrRezoo/CarApp/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{

	// OTP
	service_errors.OTPExists:   409,
	service_errors.OTPUsed:     409,
	service_errors.OTONotValid: 400,

	// User
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
