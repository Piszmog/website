package handler

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/Piszmog/website/components/core"
	"github.com/Piszmog/website/components/pages"
	"github.com/Piszmog/website/data"
)

var experienceHTML []byte
var experienceOnce sync.Once

func (h *Handler) Experience(w http.ResponseWriter, r *http.Request) {
	experienceOnce.Do(func() {
		var buf bytes.Buffer
		core.HTML("experience", pages.Experience(data.ExperienceData)).Render(r.Context(), &buf)
		experienceHTML = buf.Bytes()
	})
	h.htmlStatic(w, http.StatusOK, experienceHTML)
}
