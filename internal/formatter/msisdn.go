package formatter

import (
	"fmt"
	"math/rand"
	"time"
)

const maxMsisdnLength = 15

type msisdnFormatter struct {
	msisdnLength        int
	numberingPlanPrefix []string
	rand                *rand.Rand
}

func NewMsisdnFormatter(numberingPlansPrefix []string, maxLength int) Formatter {
	if maxLength == 0 {
		maxLength = maxMsisdnLength
	}
	if numberingPlansPrefix == nil {
		numberingPlansPrefix = make([]string, 0)
	}
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &msisdnFormatter{
		msisdnLength:        maxLength,
		numberingPlanPrefix: numberingPlansPrefix,
		rand:                seededRand,
	}
}

func (m *msisdnFormatter) Generate() string {
	if len(m.numberingPlanPrefix) == 0 {
		return generateSuffix(m.rand, maxLength)
	}
	prefix := m.numberingPlanPrefix[m.rand.Intn(len(m.numberingPlanPrefix))]

	suffix := generateSuffix(m.rand, maxLength-len(prefix))
	return fmt.Sprintf("%s%s", prefix, suffix)
}
