package handler

import (
	"net/http"
)

func (h *Handler) Resume(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/assets/img/resume.pdf", http.StatusFound)
}
