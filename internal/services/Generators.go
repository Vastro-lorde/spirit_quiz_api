package services

import (
	"math/rand"
	"strconv"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate random numbers
func GenerateRandomNumbers(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(seededRand.Intn(10))
	}
	return result
}


// generate random Alphanumeric string
func GenerateRandomAlphaNumericString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}