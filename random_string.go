package randomstring

import (
	"math/rand"
	"time"
)

// charset
var (
	Alpha       = AlphaLower + AlphaUpper
	AlphaNumber = Alpha + Number

	AlphaLower = "abcdefghijklmnopqrstuvwxyz"
	AlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Number = "1234567890"
)

// RandomAlpha generates a random string containing only alpha
func RandomAlpha(n int) string {
	return Random(n, Alpha)
}

// RandomAlphaNumber generates a random string containing only alpha and number
func RandomAlphaNumber(n int) string {
	return Random(n, AlphaNumber)
}

// RandomAlphaLower generates a random string containing only alpha in lower case
func RandomAlphaLower(n int) string {
	return Random(n, AlphaLower)
}

// RandomAlphaUpper generates a random string containing only alpha in upper case
func RandomAlphaUpper(n int) string {
	return Random(n, AlphaUpper)
}

// RandomNumber generates a random string containing only number
func RandomNumber(n int) string {
	return Random(n, Number)
}

// Random generates a random string containing only characters in given charset
func Random(n int, charset string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}
	return string(b)
}
