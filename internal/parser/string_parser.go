package parser

import (
	"errors"
	"fmt"
	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
	"regexp"
	"strings"
)

const formatRegex = `^\((\w+\|\w+)*\)$`

func generateStringFormat(details domain.ColumnDetails) (formatter.Formatter, error) {
	if details.Format == "" || details.Format == "*" {
		return formatter.NewStringFormatter(nil, details.Length), nil
	}

	result, err := regexp.MatchString(formatRegex, details.Format)
	if err != nil {
		return nil, err
	}

	if result {
		format := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(details.Format), "("), ")")
		values := strings.Split(format, "|")
		return formatter.NewStringFormatter(values, 0), nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown string format %s", details.Format))
}
