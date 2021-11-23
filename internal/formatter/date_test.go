package formatter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateFormatter_Generate(t *testing.T) {
	tests := []struct {
		name       string
		format     string
		wantLayout string
	}{
		{name: "Date format 2006-01-02", format: "yyyy-MM-dd", wantLayout: "2006-01-02"},
		{name: "Date format 2006-01-02T15:04:05", format: "yyyy-MM-ddThh:mm:ss", wantLayout: "2006-01-02T15:04:05"},
		{name: "Date format 01/02/2006", format: "MM/dd/yyyy", wantLayout: "01/02/2006"},
		{name: "Date format Jan-02-06 15:04:05", format: "MMM-dd-yy hh:mm:ss", wantLayout: "Jan-02-06 15:04:05"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewDateFormatter(tc.format).Generate()
			dd, err := time.Parse(tc.wantLayout, got)
			assert.Nil(t, err)
			now := time.Now()
			assert.Equal(t, now.Day(), dd.Day())
			assert.Equal(t, now.Month(), dd.Month())
			assert.Equal(t, now.Day(), dd.Day())
		})
	}
}
