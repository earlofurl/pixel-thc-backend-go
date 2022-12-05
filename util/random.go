package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer in the range [min, max]
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	n := rand.Intn(10) + 1
	return RandomString(n)
}

// RandomMoney returns a random money amount
func RandomMoney() int64 {
	return RandomInt(0, 1000000)
}

// RandomCurrency returns a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "JPY", "GBP", "AUD", "CAD", "CHF", "CNY", "SEK", "NZD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
