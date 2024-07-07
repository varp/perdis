package parser

import "log/slog"

type Parser struct {
	logger *slog.Logger
	sm     *stateMachine
}

func NewParser(logger *slog.Logger) *Parser {
	return &Parser{
		logger: logger,
		sm:     newStateMachine(),
	}
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
