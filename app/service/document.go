package service

import (
	"context"
	"errors"
	"io"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/raitonoberu/personal-best/app/model"
	"github.com/raitonoberu/personal-best/db/sqlc"
)

const documentExpirationOffset = time.Hour * 24 * 14
const documentLinkLifetime = time.Hour * 3

var documentBucket = os.Getenv("DOCUMENT_BUCKET")

func (s Service) SaveDocument(ctx context.Context, userID int64, reader io.Reader, name string, size int64, contentType string) error {
	filename := uuid.New().String()
	expiresAt := time.Now().Add(documentExpirationOffset)

	_, err := s.s3.PutObject(
		ctx,
		documentBucket,
		filename,
		reader,
		size,
		minio.PutObjectOptions{
			UserTags: map[string]string{
				"userID": strconv.FormatInt(userID, 10),
			},
			ContentType: contentType,
			Expires:     expiresAt,
		},
	)
	if err != nil {
		return err
	}

	return s.queries.CreateDocument(ctx, sqlc.CreateDocumentParams{
		PlayerID:  userID,
		Name:      name,
		Url:       filename,
		ExpiresAt: expiresAt,
	})
}

func (s Service) ListDocuments(ctx context.Context, userID int64) ([]model.Document, error) {
	docs, err := s.queries.ListDocuments(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]model.Document, len(docs))
	for i, d := range docs {
		result[i] = model.Document{
			ID:        d.ID,
			Name:      d.Name,
			CreatedAt: d.CreatedAt,
			ExpiresAt: d.ExpiresAt,
		}
	}
	return result, nil
}

func (s Service) GetDocument(ctx context.Context, id, userID int64, admin bool) (string, error) {
	record, err := s.queries.GetDocument(ctx, id)
	if err != nil {
		return "", err
	}
	if record.PlayerID != userID && !admin {
		return "", errors.New("Нет прав на просмотр документа")
	}

	url, err := s.s3.PresignedGetObject(ctx, documentBucket, record.Url, documentLinkLifetime, url.Values{})
	if err != nil {
		return "", err
	}
	return url.String(), err
}
