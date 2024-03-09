package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

// Handler handles requests.
type Handler struct {
	Logger *slog.Logger
}

func (h *Handler) html(ctx context.Context, w http.ResponseWriter, status int, t templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	_ = t.Render(ctx, w)
}

func (h *Handler) htmlStatic(w http.ResponseWriter, status int, html []byte) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	w.Write(html)
}
