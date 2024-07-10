package common

import (
	"log"
	"regexp"
)

const MobilePattern = `^(\+98|0)?9\d{9}$`

func ValidateMobile(mobileNumber string) bool {
	res, err := regexp.MatchString(MobilePattern, mobileNumber)
	if err != nil {
		log.Println(err.Error())
	}
	return res
}
