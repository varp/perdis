package memory

import (
	"errors"
	"log/slog"
)

type Engine struct {
	data   map[string]string
	logger *slog.Logger
}

func NewEngine(logger *slog.Logger) (*Engine, error) {
	if logger == nil {
		return nil, errors.New("nil logger was passed")
	}

	return &Engine{
		data:   make(map[string]string),
		logger: logger,
	}, nil
}

func (e *Engine) Set(key, value string) {
	e.logger.Debug("SET", "key", key, "value", value)
	e.data[key] = value
}

func (e *Engine) Get(key string) string {
	value := e.data[key]
	e.logger.Debug("GET", "key", key, "value", value)
	return value
}

func (e *Engine) Del(key string) {
	e.logger.Debug("DEL", "key", key)
	delete(e.data, key)
}
