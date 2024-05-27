package gen

import (
	"math"
	"strings"
)

func CalculateEntropy(password string) float64 {
	var charsetSize float64

	charsets := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"!@#$%^&*()-_=+{}[]|:;<>,.?/~`",
	}

	for _, charset := range charsets {
		if strings.ContainsAny(password, charset) {
			charsetSize += float64(len(charset))
		}
	}

	return math.Log2(charsetSize) * float64(len(password))
}
