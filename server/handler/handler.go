package handler

import (
	"log/slog"
	"net/http"
)

// Handler handles requests.
type Handler struct {
	Logger *slog.Logger
}

func (h *Handler) htmlStatic(w http.ResponseWriter, status int, html []byte) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(html); err != nil {
		h.Logger.Error("Failed to write response", "error", err)
	}
}
