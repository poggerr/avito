package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/poggerr/avito/internal/app"
	"github.com/poggerr/avito/internal/logger"
)

func Router(app *app.App) chi.Router {
	r := chi.NewRouter()
	r.Use(logger.WithLogging)
	r.Post("/api/segment/create", app.CreateSegment)
	r.Post("/api/segment/delete", app.DeleteSegment)
	r.Get("/api/segment/{id}", app.SegmentByUserID)
	r.Post("/api/user/segment", app.SegmentsToUser)
	r.Post("/api/segment/csv/{id}", app.CreateCSVSegment)
	return r
}
