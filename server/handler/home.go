package handler

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/Piszmog/website/components/core"
	"github.com/Piszmog/website/components/pages"
)

var homeHTML []byte
var homeOnce sync.Once

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	homeOnce.Do(func() {
		var buf bytes.Buffer
		core.HTML("", pages.Home()).Render(r.Context(), &buf)
		homeHTML = buf.Bytes()
	})
	h.htmlStatic(w, http.StatusOK, homeHTML)
}
