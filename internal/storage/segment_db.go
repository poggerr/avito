package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/models"
	"time"
)

func (strg *Storage) CreateSegmentDB(segment *models.Segment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := uuid.New()

	_, err := strg.DB.ExecContext(
		ctx,
		"INSERT INTO segments (id, slug) VALUES ($1, $2)",
		id, segment.Slug)
	if err != nil {
		return err
	}
	return nil
}

func (strg *Storage) DuplicateSegment(segment *models.Segment) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id uuid.UUID

	ans := strg.DB.QueryRowContext(ctx, "SELECT id FROM segments WHERE slug=$1", segment.Slug)
	errScan := ans.Scan(&id)
	if errScan != nil {
		logger.Initialize().Info(errScan)
		return false
	}
	return true
}

func (strg *Storage) DeleteSegmentDB(segment *models.Segment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := strg.DB.ExecContext(ctx, "DELETE FROM segments WHERE slug=$1", segment.Slug)
	if err != nil {
		return err
	}
	return nil
}

func (strg *Storage) SegmentByUserID(id *uuid.UUID) *models.Storage {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := strg.DB.QueryContext(ctx, "SELECT * FROM user_segments WHERE user_id=$1", id)
	if err != nil {
		logger.Initialize().Info(err)
	}

	storage := make(models.Storage, 0)
	for rows.Next() {
		var newID int
		var segment models.Segment
		var user uuid.UUID
		if err = rows.Scan(&newID, &user, &segment.Slug); err != nil {
			logger.Initialize().Info(err)
		}
		storage = append(storage, segment)
	}

	if err = rows.Err(); err != nil {
		logger.Initialize().Info(err)
	}

	return &storage
}

func (strg *Storage) AddSegmentToUser(ans *models.CRUDSegmentToUser, insertSlice []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := strg.DB.Begin()
	if err != nil {
		return err
	}
	for _, slug := range insertSlice {
		_, err = tx.ExecContext(
			ctx,
			"INSERT INTO user_segments (user_id, segment_slug) VALUES ($1, $2)",
			ans.UserID, slug)
		if err != nil {
			logger.Initialize().Info(err)
		}
		dateNow := time.Now()

		_, err = tx.ExecContext(
			ctx,
			"INSERT INTO segment_history (user_id, segment_slug, operation, datetime) VALUES ($1, $2, $3, $4)",
			ans.UserID, slug, "CREATE", dateNow.Format("02-2006"))
		if err != nil {
			logger.Initialize().Info(err)
		}
	}
	tx.Commit()
	return nil
}

func (strg *Storage) DeleteSegmentUser(ans *models.CRUDSegmentToUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := strg.DB.Begin()
	if err != nil {
		return err
	}
	for _, slug := range ans.Delete {
		_, err = tx.ExecContext(
			ctx,
			"DELETE FROM user_segments WHERE user_id=$1 AND segment_slug=$2",
			ans.UserID, slug)
		if err != nil {
			logger.Initialize().Info(err)
		}

		dateNow := time.Now()

		_, err = tx.ExecContext(
			ctx,
			"INSERT INTO segment_history (user_id, segment_slug, operation, datetime) VALUES ($1, $2, $3, $4)",
			ans.UserID, slug, "DELETE", dateNow.Format("02-2006"))
		if err != nil {
			logger.Initialize().Info(err)
		}
	}
	tx.Commit()
	return nil
}

func (strg *Storage) GetInfoSegmentsByUser(duration *models.CSVRequest, userID *uuid.UUID) (*models.StorageCSV, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := strg.DB.QueryContext(ctx, "SELECT * FROM segment_history WHERE user_id=$1 AND datetime=$2", userID, duration.Period)
	if err != nil {
		logger.Initialize().Info(err)
		return nil, err
	}

	storage := make(models.StorageCSV, 0)
	for rows.Next() {
		var slug string
		var operation string
		var datetime string
		if err = rows.Scan(&userID, &slug, &operation, &datetime); err != nil {
			logger.Initialize().Info(err)
			return nil, err
		}
		segment := []string{userID.String(), slug, operation, datetime}
		storage = append(storage, segment)
	}

	if err = rows.Err(); err != nil {
		logger.Initialize().Info(err)
		return nil, err
	}
	return &storage, nil
}

func (strg *Storage) CreateInsertSlice(ans *models.CRUDSegmentToUser) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var insertSlice []string
	for _, slug := range ans.Add {
		var id int
		answer := strg.DB.QueryRowContext(ctx, "SELECT id FROM user_segments WHERE user_id=$1 AND segment_slug=$2", ans.UserID, slug)
		errScan := answer.Scan(&id)
		if errScan != nil {
			logger.Initialize().Info(errScan)
			insertSlice = append(insertSlice, slug)
		}
	}
	return insertSlice
}
