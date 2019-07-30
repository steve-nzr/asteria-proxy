package server

import (
	"log"
	"os"
	"strconv"
)

var bufferSize = getBufferSize()

func getBufferSize() int64 {
	len := os.Getenv("SERVER_MAX_DATA_LENGTH")
	length, err := strconv.ParseInt(len, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return length
}
