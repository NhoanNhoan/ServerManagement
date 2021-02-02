package database

import (
	"math/rand"
	"strings"
)

func GeneratePrimaryKey(useLowerCase bool,
					useUpperCase bool,
					useNumbers bool,
					useSpecial bool,
					prefix string,
					primaryKeySize int) string {
	const LOWER_CASE string = "abcdefghijklmnopqursuvwxyz"
	const UPPER_CASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const NUMBERS string =  "123456789"
	const SPECIALS string = "!@£$%^&*()#€"

	var primaryKey strings.Builder
	charSet := ""
	if useLowerCase {
		charSet += LOWER_CASE
	}

	if useUpperCase {
		charSet += UPPER_CASE
	}

	if useNumbers {
		charSet += NUMBERS
	}

	if useSpecial {
		charSet += SPECIALS
	}

	primaryKey.WriteString(prefix)
	primaryKeySize -= len(prefix)

	for i := 0; i < primaryKeySize; i++ {
		randNum := rand.Intn(len(charSet) - 1)
		primaryKey.WriteString(string(charSet[randNum]))
	}

	return primaryKey.String()
}