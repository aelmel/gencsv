package parser

import "strings"

func trimFormatParentheses(format string) string {
	return  strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(format), "("), ")")
}