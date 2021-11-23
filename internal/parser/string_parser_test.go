package parser

import (
	"github.com/aelmel/gencsv/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse_StringParser(t *testing.T) {
	tests := map[string]struct {
		want    []string
		details domain.ColumnDetails
	}{
		"Empty string format": {
			details: domain.ColumnDetails{
				Type:   "string",
				Length: 0,
				Format: "",
			},
			want: []string{""}},

		"Pattern string": {
			details: domain.ColumnDetails{
				Type:   "string",
				Length: 10,
				Format: "(true|false)",
			},
			want: []string{"true", "false"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			formatter, err := generateStringFormat(tc.details)
			assert.Nil(t, err)
			got := formatter.Generate()
			assert.Contains(t, tc.want, got)

		})
	}
}
