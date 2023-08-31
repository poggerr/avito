package app

import (
	"encoding/json"
	"github.com/poggerr/avito/internal/models"
	"io"
	"net/http"
)

func (a *App) DeleteSegment(res http.ResponseWriter, req *http.Request) {
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

	err = a.strg.DeleteSegmentDB(&segment)
	if err != nil {
		a.sugaredLogger.Info(err)
		res.WriteHeader(http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusOK)
}
