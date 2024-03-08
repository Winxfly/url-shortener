package slogdiscard

import (
	"context"
	"golang.org/x/exp/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct {
}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// ignore log entry
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// return handler, because no attributes to save
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// return handler, because no group to save
	return h
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// always return false, because write to
	return false
}
