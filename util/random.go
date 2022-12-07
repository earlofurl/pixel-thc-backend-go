package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat generates a random decimal
func RandomFloat() float64 {
	return rand.Float64()
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomPhone generates a random phone number
func RandomPhone() string {
	return RandomString(10)
}

// RandomRole generates a random role
func RandomRole() string {
	roles := []string{"Admin", "User"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomCategory generates a random category
func RandomCategory() string {
	return RandomString(6)
}

// RandomQuantity generates a random decimal quantity
func RandomQuantity() float64 {
	return RandomFloat()
}

// RandomBool generates a random boolean value
func RandomBool() bool {
	return rand.Intn(2) == 0
}

// RandomPercent generates a random decimal percent
func RandomPercent() float64 {
	return RandomFloat()
}

// RandomDecimalTimes100 generates a random decimal times 100
func RandomDecimalTimes100() float64 {
	return RandomFloat() * 100
}

func RandomTagNumber() string {
	return RandomString(24)
}
