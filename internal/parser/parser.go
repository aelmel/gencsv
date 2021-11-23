package parser

import "github.com/aelmel/gencsv/internal/domain"

func Parse(columnDetails domain.ColumnDetails) {
	switch columnDetails.Type {
	case domain.String:
		generateStringFormat(columnDetails)
	case domain.Date:
		generateDateFormat(columnDetails)
	case domain.Imsi:
		generateImsiFormatter(columnDetails)
	case domain.Msisdn:

	}
}
