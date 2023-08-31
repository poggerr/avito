package csv_writer

import (
	"encoding/csv"
	"github.com/google/uuid"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/models"
	"github.com/poggerr/avito/internal/storage"
	"os"
)

func CreateCSVFile(duration *models.CSVRequest, userID *uuid.UUID, strg *storage.Storage) (string, error) {
	storageCSV, err := strg.GetInfoSegmentsByUser(duration, userID)
	if err != nil {
		logger.Initialize().Info(err)
		return "", err
	}

	filename := "files/" + userID.String() + ".csv"

	file, err := os.Create(filename)
	if err != nil {
		logger.Initialize().Info(err)
		return "", err
	}
	defer file.Close()
	w := csv.NewWriter(file)

	for _, record := range *storageCSV {
		if err = w.Write(record); err != nil {
			logger.Initialize().Info(err)
			return "", err
		}
	}
	w.Flush()
	if err = w.Error(); err != nil {
		logger.Initialize().Info(err)
		return "", err
	}
	return filename, nil
}
