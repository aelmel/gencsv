package formatter

import (
	"fmt"
	"math/rand"
	"time"
)

const maxLength = 15

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
		return generateSuffix(imsi.rand, maxLength)
	}
	prefix := imsi.prefixes[imsi.rand.Intn(len(imsi.prefixes))]

	suffix := generateSuffix(imsi.rand, maxLength-len(prefix))
	return fmt.Sprintf("%s%s", prefix, suffix)
}
