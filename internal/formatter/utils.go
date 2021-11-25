package formatter

import "math/rand"

const digits = "0123456789"



func generateSuffix(rand *rand.Rand, length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}
