package service

import (
	"database/sql"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

type Service struct {
	db      *sql.DB
	queries *sqlc.Queries
	s3      *minio.Client
}

func New(db *sql.DB) (*Service, error) {
	queries := sqlc.New(db)

	s3endpoint := os.Getenv("MINIO_ENDPOINT")
	s3user := os.Getenv("MINIO_ROOT_USER")
	s3password := os.Getenv("MINIO_ROOT_PASSWORD")
	s3ssl := os.Getenv("MINIO_SECURE") == "true"

	s3, err := minio.New(s3endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3user, s3password, ""),
		Secure: s3ssl,
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		db:      db,
		queries: queries,
		s3:      s3,
	}, nil
}
