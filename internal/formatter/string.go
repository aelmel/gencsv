package formatter

import (
	"math/rand"
	"time"
)

type Formatter interface {
	Generate() string
}

type stringFormatter struct {
	values []string
	length int
	rand   *rand.Rand
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewStringFormatter(values []string, length int) Formatter {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &stringFormatter{
		values: values,
		length: length,
		rand:   seededRand,
	}
}

func (s *stringFormatter) Generate() string {
	if s.values != nil || len(s.values) != 0 {
		return s.values[s.rand.Intn(len(s.values))]
	}
	return s.randomString()

}

func (s *stringFormatter) randomString() string {
	b := make([]byte, s.length)

	for i := range b {
		b[i] = charset[s.rand.Intn(len(charset))]
	}
	return string(b)
}
