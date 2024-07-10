package parser

import (
	"errors"
	"log/slog"
)

type Parser struct {
	logger *slog.Logger
	sm     *stateMachine
}

func NewParser(logger *slog.Logger) (*Parser, error) {
	if logger == nil {
		return nil, errors.New("non nil logger required")
	}

	return &Parser{
		logger: logger,
		sm:     newStateMachine(),
	}, nil
}

func (p *Parser) ParseQuery(query string) ([]string, error) {
	tokens, err := p.sm.parse(query)
	if err != nil {
		return nil, err
	}

	p.logger.Debug(
		"parsed query",
		slog.String("query", query),
		slog.Any("tokens", tokens),
	)

	return tokens, nil
}
