package app

import (
	"encoding/json"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/models"
	"io"
	"net/http"
)

func (a *App) CreateSegment(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var segment models.Segment

	err = json.Unmarshal(body, &segment)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	isDuplicate := a.strg.DuplicateSegment(&segment)
	if isDuplicate != false {
		res.WriteHeader(http.StatusConflict)
		return
	}

	err = a.strg.CreateSegmentDB(&segment)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)
}

func (a *App) SegmentsToUser(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var ans models.CRUDSegmentToUser

	err = json.Unmarshal(body, &ans)
	if err != nil {
		logger.Initialize().Info(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	insertSlice := a.strg.CreateInsertSlice(&ans)

	err = a.strg.AddSegmentToUser(&ans, insertSlice)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.strg.DeleteSegmentUser(&ans)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
