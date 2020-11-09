package rule

import (
	"regexp"
	"strconv"
	"unicode"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsAlphaOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsSpaceExist(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

func IsAlphaNumeric(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return re.MatchString(s)
}

func IsUppercaseLetterExist(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func IsLowercaseLetterExist(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}
