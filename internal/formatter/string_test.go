package formatter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringFormatter_GenerateRandom(t *testing.T) {
	formatter := NewStringFormatter(nil, 12)

	for i := 0; i < 100; i++ {
		require.NotEqual(t, formatter.Generate(), formatter.Generate())
	}
}

func TestStringFormatter_GenerateOneValue(t *testing.T) {
	formatter := NewStringFormatter([]string{"test"}, 0)

	for i := 0; i < 100; i++ {
		require.Equal(t, formatter.Generate(), formatter.Generate())
	}
}

func TestStringFormatter_GenerateValues(t *testing.T) {
	values := []string{"one", "two", "three", "four", "five"}
	formatter := NewStringFormatter(values, 0)
	for i := 0; i < 100; i++ {
		response := formatter.Generate()
		require.True(t, contains(values, response), fmt.Sprintf("Original slice doesn't contain %s", response))
	}
}

func TestStringFormatter_GenerateZeroLength(t *testing.T) {
	formatter := NewStringFormatter(nil, 0)
	fmt.Println(formatter.Generate())
}

func contains(values []string, key string) bool {
	for _, value := range values {
		if value == key {
			return true
		}
	}
	return false
}
