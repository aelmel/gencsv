package main

import (
	"flag"
	"github.com/aelmel/gencsv/internal/processor"
	"go.uber.org/zap"
	"os"
)

func main() {
	var (
		inputFile  = flag.String("input", "D:/gen.json", "Input file format")
		outputFile = flag.String("output-dir", "./", "Output file directory")
	)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if *inputFile == "" {
		sugar.Errorf("No input file provided")
		os.Exit(1)
	}

	processor, err := processor.NewFileProcessor(*inputFile, sugar)
	if err != nil {
		os.Exit(1)
	}

	processor.GenerateOutput(*outputFile)
}
