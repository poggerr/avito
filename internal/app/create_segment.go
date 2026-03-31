package app

import (
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/models"
	"net/http"
	"strings"
)

func (a *App) CreateSegment(res http.ResponseWriter, req *http.Request) {
	var segment models.Segment
	if err := decodeJSONBody(req, &segment); err != nil {
		a.sugaredLogger.Infow("invalid create segment request", "error", err)
		writeError(res, http.StatusBadRequest, "invalid JSON body")
		return
	}
	segment.Slug = strings.TrimSpace(segment.Slug)
	if segment.Slug == "" {
		writeError(res, http.StatusBadRequest, "segment is required")
		return
	}

	isDuplicate := a.strg.DuplicateSegment(&segment)
	if isDuplicate {
		writeError(res, http.StatusConflict, "segment already exists")
		return
	}

	if err := a.strg.CreateSegmentDB(&segment); err != nil {
		a.sugaredLogger.Infow("failed to create segment", "segment", segment.Slug, "error", err)
		writeError(res, http.StatusBadRequest, "failed to create segment")
		return
	}

	writeJSON(res, http.StatusCreated, map[string]string{"status": "created"})
}

func (a *App) SegmentsToUser(res http.ResponseWriter, req *http.Request) {
	var ans models.CRUDSegmentToUser
	if err := decodeJSONBody(req, &ans); err != nil {
		a.sugaredLogger.Infow("invalid update user segments request", "error", err)
		writeError(res, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if ans.UserID == uuid.Nil {
		writeError(res, http.StatusBadRequest, "user is required")
		return
	}

	insertSlice := a.strg.CreateInsertSlice(&ans)

	if err := a.strg.AddSegmentToUser(&ans, insertSlice); err != nil {
		a.sugaredLogger.Infow("failed to add user segments", "user", ans.UserID, "error", err)
		writeError(res, http.StatusBadRequest, "failed to add user segments")
		return
	}

	if err := a.strg.DeleteSegmentUser(&ans); err != nil {
		a.sugaredLogger.Infow("failed to delete user segments", "user", ans.UserID, "error", err)
		writeError(res, http.StatusBadRequest, "failed to delete user segments")
		return
	}

	writeJSON(res, http.StatusCreated, map[string]string{"status": "updated"})
}
