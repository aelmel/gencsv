package formatter

type msisdnFormatter struct{}

func NewMsisdnFormatter() Formatter {
	return &msisdnFormatter{}
}

func (m msisdnFormatter) Generate() string {
	return "msisdn"
}
