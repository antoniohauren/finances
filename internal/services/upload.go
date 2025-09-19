package services

import (
	"io"

	"github.com/antoniohauren/finances/internal/storage"
	"github.com/google/uuid"
)

func UploadReceipt(file io.Reader, orignalName string, userID uuid.UUID) (string, error) {
	res, err := storage.UploadFile(file, orignalName, "receipts", userID.String())

	if err != nil {
		return "", err
	}

	return storage.GetFileURL(res.BucketName, res.Key)
}

func UploadProfilePicture(file io.Reader, orignalName string, userID uuid.UUID) (string, error) {
	res, err := storage.UploadFile(file, orignalName, "images", userID.String())

	if err != nil {
		return "", err
	}

	return storage.GetPublicFileURL(res.BucketName, res.Key)
}
