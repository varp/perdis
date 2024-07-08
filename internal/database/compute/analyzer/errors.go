package analyzer

import "errors"

var (
	errAnalyzerInvalidCommand = errors.New("invalid command")
	errAnalyzerInvalidQuery   = errors.New("invalid query")
	errAnalyzerUnknownCommand = errors.New("unknown command")
)
