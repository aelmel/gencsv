package parser

import "gencsv/internal/domain"

func Parse(columnDetails domain.ColumnDetails) {
	switch columnDetails.Type {
	case domain.String:
		generateStringFormat(columnDetails)
	case domain.Date:
	case domain.Imsi:
	case domain.Msisdn:

	}
}

