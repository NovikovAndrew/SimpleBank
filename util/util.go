package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generator integers between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generator string of size n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate the string if size 10
func RandomOwner() string {
	return RandomString(10)
}

// RandomMoney generate random money between 0 and 1000
func RandomMoney() int {
	return int(RandomInt(0, 1000))
}

// RandomCurrency generate random currency
func RandomCurrency() string {
	currencies := []string{"UER", "USD", "RUB", "KZT"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
