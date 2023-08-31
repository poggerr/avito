package models

import (
	"github.com/google/uuid"
)

type Segment struct {
	Slug string `json:"segment"`
}

type User struct {
	Username string `json:"username"`
}

type Storage []Segment

type CRUDSegmentToUser struct {
	Add    []string  `json:"add"`
	Delete []string  `json:"delete"`
	UserID uuid.UUID `json:"user"`
}

type CSVRequest struct {
	Period string `json:"period"`
}

type StorageCSV [][]string

type CSVLink struct {
	Link string `json:"link"`
}
