package handler

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/Piszmog/website/components/core"
	"github.com/Piszmog/website/components/pages"
)

var aboutHTML []byte
var aboutOnce sync.Once

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	aboutOnce.Do(func() {
		var buf bytes.Buffer
		_ = core.HTML("about", pages.About()).Render(r.Context(), &buf)
		aboutHTML = buf.Bytes()
	})
	h.htmlStatic(w, http.StatusOK, aboutHTML)
}
