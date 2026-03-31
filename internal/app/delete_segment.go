package app

import (
	"github.com/poggerr/avito/internal/models"
	"net/http"
	"strings"
)

func (a *App) DeleteSegment(res http.ResponseWriter, req *http.Request) {
	var segment models.Segment
	if err := decodeJSONBody(req, &segment); err != nil {
		a.sugaredLogger.Infow("invalid delete segment request", "error", err)
		writeError(res, http.StatusBadRequest, "invalid JSON body")
		return
	}
	segment.Slug = strings.TrimSpace(segment.Slug)
	if segment.Slug == "" {
		writeError(res, http.StatusBadRequest, "segment is required")
		return
	}

	if err := a.strg.DeleteSegmentDB(&segment); err != nil {
		a.sugaredLogger.Infow("failed to delete segment", "segment", segment.Slug, "error", err)
		writeError(res, http.StatusBadRequest, "failed to delete segment")
		return
	}

	writeJSON(res, http.StatusOK, map[string]string{"status": "deleted"})
}
