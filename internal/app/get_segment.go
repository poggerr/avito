package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func (a *App) SegmentByUserID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	newID, err := uuid.Parse(id)
	if err != nil {
		writeError(res, http.StatusBadRequest, "invalid user id")
		return
	}
	storage := a.strg.SegmentByUserID(&newID)
	writeJSON(res, http.StatusOK, storage)
}
