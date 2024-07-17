package test

import (
	"context"
	"log/slog"
)

type nopLoggerHandler struct {
}

func NewNopLoggerHandler() slog.Handler {
	return &nopLoggerHandler{}
}

func (n *nopLoggerHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (n *nopLoggerHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (n *nopLoggerHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return n
}

func (n *nopLoggerHandler) WithGroup(_ string) slog.Handler {
	return n
}

func NewNopLogger() *slog.Logger {
	return slog.New(NewNopLoggerHandler())
}
