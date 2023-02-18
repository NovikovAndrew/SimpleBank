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

// RandomInt generate random integer value between min and max
func RandomInt(min, max int) int64 {
	return int64(min + rand.Intn(max-min+1))
}

// GenerateRandonString generate random string by len of b
func GenerateRandonString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return GenerateRandonString(6)
}

func GenerateRandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "RUB", "TEN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
