package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/models"
	"time"
)

func (strg *Storage) DuplicateUser(user *models.User) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id uuid.UUID

	ans := strg.DB.QueryRowContext(ctx, "SELECT id FROM users WHERE username=$1", user.Username)
	errScan := ans.Scan(&id)
	if errScan != nil {
		logger.Initialize().Info(errScan)
		return false
	}
	return true
}

func (strg *Storage) VerifyUser(userID *uuid.UUID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User

	ans := strg.DB.QueryRowContext(ctx, "SELECT username FROM users WHERE id=$1", userID)
	errScan := ans.Scan(&user.Username)
	if errScan != nil {
		logger.Initialize().Info(errScan)
		return false
	}
	return true
}
