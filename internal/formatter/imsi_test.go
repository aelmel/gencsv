package formatter

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestImsiFormatter_Generate(t *testing.T) {
	t.Run("Test random imsi", func(t *testing.T) {
		assert.Len(t, NewImsiFormatter(make([]string, 0)).Generate(), 15)
	})
	t.Run("test based on imsi pattern ", func(t *testing.T) {
		prefixes := []string{"90"}
		imsi := make([]string, 20)
		for i := 0; i < 20; i++ {
			imsi = append(imsi, NewImsiFormatter(prefixes).Generate())
		}

		for _, i := range imsi {
			strings.HasPrefix(i, prefixes[0])
		}
	})
}
