package processor

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"sync"

	"go.uber.org/zap"

	"github.com/aelmel/gencsv/internal/domain"
	"github.com/aelmel/gencsv/internal/formatter"
	"github.com/aelmel/gencsv/internal/parser"
)

type fileProcessor struct {
	formatters    []formatter.Formatter
	fileStructure domain.FileStructure
	logger        *zap.SugaredLogger
	rows          chan string
	wg            *sync.WaitGroup
}

func NewFileProcessor(input string, logger *zap.SugaredLogger) (*fileProcessor, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		logger.Error("Error reading input file.", zap.String("file", input), zap.String("err", err.Error()))
		return nil, err
	}

	var fStructure domain.FileStructure
	err = json.Unmarshal(data, &fStructure)
	if err != nil {
		logger.Error("Error parsing input file", zap.String("file", input), zap.String("err", err.Error()))
		return nil, err
	}

	formatters, err := createFormatters(fStructure)
	if err != nil {
		logger.Error("Error parsing input file", zap.String("file", input), zap.String("err", err.Error()))
		return nil, err
	}

	wg := &sync.WaitGroup{}
	rowChan := make(chan string)
	return &fileProcessor{formatters: formatters, fileStructure: fStructure, logger: logger, wg: wg, rows: rowChan}, nil

}

func (fp *fileProcessor) GenerateOutput(output string) error {
	f, err := os.OpenFile(path.Join(output, fp.fileStructure.Filename), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fp.logger.Error("error opening file", zap.Error(err))
		return err
	}
	defer f.Close()

	for i := 0; i < fp.fileStructure.Rows; i++ {
		fp.wg.Add(1)
		go fp.generateRow()
	}

	go fp.monitor()
	for row := range fp.rows {
		f.WriteString(row + "\n")
	}
	return nil
}

func createFormatters(fStructure domain.FileStructure) ([]formatter.Formatter, error) {
	formatters := make([]formatter.Formatter, len(fStructure.Columns))
	for i, details := range fStructure.Columns {
		columnFormatter, err := parser.Parse(details)
		if err != nil {
			return nil, err
		}
		formatters[i] = columnFormatter
	}
	return formatters, nil
}

func (fp *fileProcessor) monitor() {
	fp.wg.Wait()
	close(fp.rows)
}

func (fp *fileProcessor) generateRow() {
	defer fp.wg.Done()
	row := make([]string, len(fp.formatters))
	for i, columnFormatter := range fp.formatters {
		row[i] = columnFormatter.Generate()
	}

	fp.rows <- strings.Join(row, ",")
}
