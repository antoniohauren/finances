package storage

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

var s3Client *s3.Client

func NewS3Client() (*s3.Client, error) {
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")
	region := os.Getenv("S3_REGION")

	if region == "" {
		region = "us-east-1"
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))

	if err != nil {
		return nil, fmt.Errorf("loading aws config: %w", err)
	}

	creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	s3Client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = region
		o.Credentials = creds
		o.UsePathStyle = true
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	})

	slog.Info("connected to s3 client!")
	return s3Client, nil
}

type UploadFileReturn struct {
	BucketName string
	Key        string
}

func UploadFile(file io.Reader, orignalName string, bucketName string, uploadDir string) (*UploadFileReturn, error) {
	id := uuid.New()
	ext := filepath.Ext(orignalName)
	fileKey := filepath.Join(uploadDir, id.String()+ext)

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileKey),
		Body:   file,
	})

	if err != nil {
		slog.Error("upload", "error", err)
		return nil, err
	}

	return &UploadFileReturn{
		BucketName: bucketName,
		Key:        fileKey,
	}, nil
}

func GetFileURL(bucketName string, key string) (string, error) {
	presigner := s3.NewPresignClient(s3Client)
	presignedReq, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		slog.Error("presign-url", "error", err)
		return "", fmt.Errorf("failed to generate presigned url: %w", err)
	}

	return presignedReq.URL, nil
}
