package database

import (
	"errors"
	"fmt"
	"go.vardan.dev/perdis/internal/database/compute/analyzer"
	"go.vardan.dev/perdis/internal/database/compute/parser"
	"go.vardan.dev/perdis/internal/database/storage"
	"go.vardan.dev/perdis/internal/database/storage/memory"
	"log/slog"
)

const (
	okPrefix  = "[OK]"
	errPrefix = "[ERROR]"
)

type Database struct {
	logger *slog.Logger

	engine   storage.Engine
	parser   *parser.Parser
	analyzer *analyzer.Analyzer
}

func Start(logger *slog.Logger) (*Database, error) {
	if logger == nil {
		return nil, errors.New("non nil logger required")
	}

	var err error
	p, err := parser.NewParser(logger)
	if err != nil {
		return nil, err
	}
	a, err := analyzer.NewAnalyzer(logger)
	if err != nil {
		return nil, err
	}

	e, err := memory.NewEngine(logger)
	if err != nil {
		return nil, err
	}

	return &Database{
		logger:   logger,
		engine:   e,
		parser:   p,
		analyzer: a,
	}, nil
}

func (d *Database) Execute(query string) string {
	var err error

	tokens, err := d.parser.ParseQuery(query)
	if err != nil {
		return d.formatResult(errPrefix, "failed to parse query: "+err.Error())
	}

	compiledQuery, err := d.analyzer.Analyze(tokens)
	if err != nil {
		return d.formatResult(errPrefix, "failed to analyze query: "+err.Error())
	}

	return d.processQuery(compiledQuery)
}

func (d *Database) formatResult(prefix, message string) string {
	return fmt.Sprintf("%s %s", prefix, message)
}

func (d *Database) processQuery(query analyzer.Query) string {
	commandId := query.CommandId()
	args := query.Args()

	switch commandId {
	case analyzer.CommandGetID:
		return d.formatResult(okPrefix, d.engine.Get(args[0]))
	case analyzer.CommandSetID:
		d.engine.Set(args[0], args[1])
		return okPrefix
	case analyzer.CommandDelID:
		d.engine.Del(args[0])
		return okPrefix
	}

	return d.formatResult(errPrefix, "database is incorrect state")
}
