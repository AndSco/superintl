package utils

import (
	"log"
	"unicode"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func HasUpperCase(word string) bool {
	hasUpper := false
	for _, r := range word {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}
	return hasUpper
}
