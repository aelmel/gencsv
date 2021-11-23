package formatter

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLength = 15
const digits = "0123456789"

type imsiFormatter struct {
	prefixes []string
	rand     *rand.Rand
}

func NewImsiFormatter(prefixes []string) Formatter {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &imsiFormatter{prefixes: prefixes, rand: seededRand}
}

func (imsi *imsiFormatter) Generate() string {
	if len(imsi.prefixes) == 0 {
		return imsi.generateSuffix(maxLength)
	}
	prefix := imsi.prefixes[imsi.rand.Intn(len(imsi.prefixes))]

	suffix := imsi.generateSuffix(maxLength - len(prefix))
	return fmt.Sprintf("%s%s", prefix, suffix)
}

func (imsi *imsiFormatter) generateSuffix(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = digits[imsi.rand.Intn(len(digits))]
	}
	return string(b)
}
