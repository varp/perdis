package analyzer

import (
	"errors"
	"log/slog"
)

type Analyzer struct {
	logger *slog.Logger
}

func NewAnalyzer(logger *slog.Logger) (*Analyzer, error) {
	if logger == nil {
		return nil, errors.New("nil logger")
	}

	return &Analyzer{
		logger: logger,
	}, nil
}

func (a *Analyzer) Analyze(queryTokens []string) (Query, error) {
	if len(queryTokens) == 0 {
		a.logger.Error("empty tokens was passed")
		return Query{}, errAnalyzerInvalidQuery
	}

	commandName := queryTokens[0]
	commandArgs := queryTokens[1:]
	command := getCommandIDByName(commandName)

	if command == CommandUnknownID {
		a.logger.Error(
			"unknown command was passed",
			"command", commandName,
		)
		return Query{}, errAnalyzerUnknownCommand
	}

	switch commandName {
	case CommandDel, CommandGet:
		if len(commandArgs) == 0 || len(commandArgs) > 1 {
			a.logger.Error(
				"invalid arguments passed",
				"command", commandName,
				"args", commandArgs,
			)
			return Query{}, errAnalyzerInvalidCommand
		}
	case CommandSet:
		if len(commandArgs) < 2 {
			a.logger.Error(
				"invalid arguments passed",
				"command", commandName,
				"args", commandArgs,
			)
			return Query{}, errAnalyzerInvalidCommand
		}
	}

	return Query{commandId: command, args: commandArgs}, nil
}
