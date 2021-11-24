package parser

import "github.com/aelmel/gencsv/internal/formatter"

func generateMsisdnParser() (formatter.Formatter, error) {
	return formatter.NewMsisdnFormatter(), nil
}
