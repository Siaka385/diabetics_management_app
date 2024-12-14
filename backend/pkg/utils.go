package pkg

import (
	"os"
	"strconv"
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
