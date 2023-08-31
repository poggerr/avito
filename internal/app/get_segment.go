package app

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/logger"
	"net/http"
)

func (a *App) SegmentByUserID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	newID := uuid.Must(uuid.Parse(id))
	storage := a.strg.SegmentByUserID(&newID)

	marshal, err := json.Marshal(storage)
	if err != nil {
		logger.Initialize().Info(err)
	}

	res.Header().Set("content-type", "application/json ")
	res.WriteHeader(http.StatusOK)
	res.Write(marshal)
}
