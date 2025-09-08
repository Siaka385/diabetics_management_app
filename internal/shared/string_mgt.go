package shared

import (
	"math/rand"
	"strings"
	"time"
)

const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

/*
* GenerateRandomString generates a random string of the specified length
 */
func GenerateRandomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // Create a new random generator
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = CHARSET[r.Intn(len(CHARSET))]
	}
	return string(result)
}

func GenerateShortName(fullName string) string {
	name := strings.TrimSpace(fullName)
	words := strings.Fields(name)
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		if len(words[0]) > 0 {
			return string(words[0][0])
		}
		return ""
	}
	shortName := ""
	for _, word := range words {
		if len(word) > 0 {
			shortName += string(word[0])
		}
	}
	return shortName
}
