package utils

import (
	"regexp"
	"strings"
)

func CleanNumber(data, varName string) uint {
	number := uint(0)
	if data == "" {
		return number
	}

	re := regexp.MustCompile("[^0-9]+")
	cleanedNumber := re.ReplaceAllString(data, "")

	cleanedNumber = strings.TrimPrefix(cleanedNumber, "0")

	number = uint(ConvStrToUint(cleanedNumber, varName))

	return number
}
