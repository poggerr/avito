package app

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/csv_writer"
	"github.com/poggerr/avito/internal/models"
	"io"
	"net/http"
)

func (a *App) CreateCSVSegment(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	newID := uuid.Must(uuid.Parse(id))
	body, err := io.ReadAll(req.Body)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var segment models.CSVRequest

	err = json.Unmarshal(body, &segment)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	filename, err := csv_writer.CreateCSVFile(&segment, &newID, a.strg)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var link models.CSVLink

	link.Link = filename

	marshal, err := json.Marshal(link)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json ")
	res.WriteHeader(http.StatusOK)
	res.Write(marshal)
}
