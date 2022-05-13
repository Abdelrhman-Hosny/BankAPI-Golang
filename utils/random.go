package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]

		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMoney() int64 {
	return RandomInt(0, 100000)
}

func RandomEntry() int64 {
	return RandomInt(-1000, 1000)
}

func RandomOwner() string {
	return RandomString(10)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
