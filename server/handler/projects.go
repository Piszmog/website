package handler

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/Piszmog/website/components/core"
	"github.com/Piszmog/website/components/pages"
	"github.com/Piszmog/website/data"
)

var projectsHTML []byte
var projectsOnce sync.Once

func (h *Handler) Projects(w http.ResponseWriter, r *http.Request) {
	projectsOnce.Do(func() {
		var buf bytes.Buffer
		core.HTML("projects", pages.Projects(data.ProjectsData)).Render(r.Context(), &buf)
		projectsHTML = buf.Bytes()
	})
	h.htmlStatic(w, http.StatusOK, projectsHTML)
}
