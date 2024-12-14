package utils

import (
	"os"
	"strconv"
	"time"

	"math/rand"
)

const (
	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Port() int16 {
	var port int16 = 9000

	portStr, found := os.LookupEnv("PORT")
	if !found {
		return port
	}
	iport, err := strconv.Atoi(portStr)
	if err != nil {
		return port
	}
	return int16(iport)
}

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
