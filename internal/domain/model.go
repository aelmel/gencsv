package domain

type FileStructure struct {
	Filename     string          `json:"filename"`
	ColumnNumber int             `json:"column_number"`
	Header       string          `json:"header"`
	Rows         int             `json:"rows"`
	Separator    string          `json:"separator"`
	Columns      []ColumnDetails `json:"columns"`
}

type ColumnDetails struct {
	Type   ColumnType `json:"type"`
	Length int        `json:"length"`
	Format string     `json:"format"`
}

type ColumnType string

const (
	Imsi   ColumnType = "imsi"
	Msisdn            = "msisdn"
	Date              = "date"
	String            = "string"
	Number            = "number"
)
