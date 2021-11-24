package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
)

const imsiFormatRegex = `^\(\d+(\|\d+)*\)$`

func generateImsiFormatter(details domain.ColumnDetails) (formatter.Formatter, error) {
	if details.Format == "" || details.Format == "*" {
		return formatter.NewImsiFormatter(make([]string, 0)), nil
	}

	result, err := regexp.MatchString(imsiFormatRegex, details.Format)
	if err != nil {
		return nil, err
	}

	if result {
		format := trimFormatParentheses(details.Format)
		values := strings.Split(format, "|")
		return formatter.NewImsiFormatter(values), nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown string format %s", details.Format))
}
