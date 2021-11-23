package parser

import (
	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
)

func generateDateFormat(details domain.ColumnDetails) (formatter.Formatter, error) {
	return formatter.NewDateFormatter(details.Format), nil
}
