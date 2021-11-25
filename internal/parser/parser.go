package parser

import (
	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
)

func Parse(columnDetails domain.ColumnDetails) (formatter.Formatter, error) {
	switch columnDetails.Type {
	case domain.String:
		return generateStringFormat(columnDetails)
	case domain.Date:
		return generateDateFormat(columnDetails)
	case domain.Imsi:
		return generateImsiFormatter(columnDetails)
	case domain.Msisdn:
		return generateMsisdnParser(columnDetails)
	}
	return nil, nil
}
