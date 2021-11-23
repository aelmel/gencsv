package formatter

import (
	"strings"
	"time"
)

var datePlaceHolders = []struct {
	got  string
	want string
}{
	{"DDDD", "Monday"},
	{"dddd", "Monday"},
	{"DDD", "Mon"},
	{"ddd", "Mon"},
	{"DD", "02"},
	{"dd", "02"},
	{"D", "2"},
	{"d", "2"},
	{"yyyy", "2006"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"yy", "06"},
	{"ZZZZ", "-0700"},
	{"ZZZ", "MST"},
	{"ZZ", "Z07:00"},
	{"hh", "15"},
	{"h", "03"},
	{"mm", "04"},
	{"ss", "05"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
}

type dateFormatter struct {
	layout string
}

func NewDateFormatter(format string) Formatter {
	for _, mapping := range datePlaceHolders {
		format = strings.Replace(format, mapping.got, mapping.want, -1)
	}
	return &dateFormatter{layout: format}
}

func (d *dateFormatter) Generate() string {
	return time.Now().Format(d.layout)
}
