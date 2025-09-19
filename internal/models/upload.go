package models

import (
	"github.com/google/uuid"
)

type Upload struct {
	UserID     uuid.UUID
	BucketName string
	Key        string
}

type GetFileResponse struct {
	UserID uuid.UUID
	URL    string
}
