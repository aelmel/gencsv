package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
)

func generateMsisdnParser(details domain.ColumnDetails) (formatter.Formatter, error) {
	if details.Format == "" || details.Format == "*" {
		formatter.NewMsisdnFormatter(make([]string, 0), details.Length)
	}

	result, err := regexp.MatchString(imsiMsisdnFormatRegex, details.Format)
	if err != nil {
		return nil, err
	}

	if result {
		format := trimFormatParentheses(details.Format)
		values := strings.Split(format, "|")
		return formatter.NewMsisdnFormatter(values, details.Length), nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown msisdn format %s", details.Format))

}
