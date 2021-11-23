package parser

import (
	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
)

func generateStringFormat(details domain.ColumnDetails) formatter.Formatter {
	if details.Format == "" || details.Format == "*" {
		return formatter.NewStringFormatter(nil, details.Length)
	}
	return nil
}
