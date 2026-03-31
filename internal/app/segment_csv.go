package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/csv_writer"
	"github.com/poggerr/avito/internal/models"
	"net/http"
)

func (a *App) CreateCSVSegment(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	newID, err := uuid.Parse(id)
	if err != nil {
		writeError(res, http.StatusBadRequest, "invalid user id")
		return
	}

	var segment models.CSVRequest
	if err = decodeJSONBody(req, &segment); err != nil {
		a.sugaredLogger.Infow("invalid create csv request", "error", err)
		writeError(res, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if segment.Period == "" {
		writeError(res, http.StatusBadRequest, "period is required")
		return
	}

	filename, err := csv_writer.CreateCSVFile(&segment, &newID, a.strg)
	if err != nil {
		a.sugaredLogger.Infow("failed to create csv file", "user", newID, "period", segment.Period, "error", err)
		writeError(res, http.StatusInternalServerError, "failed to create csv file")
		return
	}

	writeJSON(res, http.StatusOK, models.CSVLink{Link: filename})
}
