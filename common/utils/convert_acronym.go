package utils

import "strings"

func GetAcronym(name string) string {
	words := strings.Fields(strings.ToLower(name))
	acronym := ""

	for _, word := range words {
		if word != "dan" {
			acronym += string(word[0])
		}
	}

	return strings.ToUpper(acronym)
}
