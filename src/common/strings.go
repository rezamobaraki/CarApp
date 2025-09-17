package common

import (
	"github.com/rezamobaraki/CarApp/config"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func GenerateOtp() string {
	cfg := config.GetConfig()
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	minNum := int(math.Pow(10, float64(cfg.OTP.Digits-1)))   // 10^d-1 100000
	maxNum := int(math.Pow(10, float64(cfg.OTP.Digits)) - 1) // 999999 = 1000000 - 1 (10^d) -1

	var num = rng.Intn(maxNum-minNum+1) + minNum
	return strconv.Itoa(num)
}

func CheckPassword(password string) bool {
	conf := config.GetConfig()

	if conf.Password.IncludeChars && !hasChars(password) {
		return false
	}

	if conf.Password.IncludeDigits && !hasDigits(password) {
		return false
	}

	if conf.Password.IncludeUppers && !hasUppers(password) {
		return false
	}

	if conf.Password.IncludeLowers && !hasLowers(password) {
		return false
	}

	if len(password) < conf.Password.MinLength || len(password) > conf.Password.MaxLength {
		return false
	}
	return true
}

func hasChars(password string) bool {
	_, err := regexp.MatchString(`[a-zA-Z]`, password)
	if err != nil {
		return false
	}
	return true
}

func hasDigits(password string) bool {
	_, err := regexp.MatchString(`[0-9]`, password)
	if err != nil {
		return false
	}
	return true
}

func hasUppers(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) && unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func hasLowers(password string) bool {
	_, err := regexp.MatchString(`[a-z]`, password)
	if err != nil {
		return false
	}
	return true
}

func PasswordGenerator() string {
	rand.Seed(time.Now().UnixNano())
	conf := config.GetConfig()

	// Ensure the password meets the minimum criteria
	minLength := conf.Password.MinLength

	// Character sets
	lowerCharSet := "abcdefghijklmnopqrstuvwxyz"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitCharSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet + digitCharSet

	// Ensuring at least one character from each set is included
	password := string(lowerCharSet[rand.Intn(len(lowerCharSet))]) +
		string(upperCharSet[rand.Intn(len(upperCharSet))]) +
		string(digitCharSet[rand.Intn(len(digitCharSet))])

	// Filling the rest of the password length with random characters from all sets
	for i := 3; i < minLength; i++ {
		password += string(allCharSet[rand.Intn(len(allCharSet))])
	}

	// Shuffle the password to avoid predictable patterns
	password = shuffleString(password)

	return password
}

// shuffleString shuffles the characters in a string
func shuffleString(str string) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	runes := []rune(str)
	for len(runes) > 0 {
		n := len(runes)
		randIndex := r.Intn(n)
		runes[n-1], runes[randIndex] = runes[randIndex], runes[n-1]
		runes = runes[:n-1]
	}
	return string(runes)
}

// To snake case : CountryId -> country_id
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
