package router

import (
	"log/slog"
	"net/http"

	"github.com/Piszmog/website/dist"
	"github.com/Piszmog/website/server/handler"
	"github.com/Piszmog/website/server/middleware"
)

func New(logger *slog.Logger) http.Handler {
	h := &handler.Handler{
		Logger: logger,
	}

	mux := http.NewServeMux()

	mux.Handle(http.MethodGet+" /assets/", middleware.CacheMiddleware(http.FileServer(http.FS(dist.AssetsDir))))
	mux.HandleFunc(http.MethodGet+" /", h.Home)
	mux.HandleFunc(http.MethodGet+" /about", h.About)
	mux.HandleFunc(http.MethodGet+" /experience", h.Experience)
	mux.HandleFunc(http.MethodGet+" /projects", h.Projects)
	mux.HandleFunc(http.MethodGet+" /resume", h.Resume)

	return middleware.NewLoggingMiddleware(logger, mux)
}
