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

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < k; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currs := []string{
		"EUR",
		"USD",
		"GBP",
	}
	k := len(currs)
	return currs[rand.Intn(k)]
}