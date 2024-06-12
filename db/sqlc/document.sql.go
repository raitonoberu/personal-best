// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: document.sql

package sqlc

import (
	"context"
	"time"
)

const createDocument = `-- name: CreateDocument :exec
INSERT INTO
    documents (player_id, name, url, expires_at)
VALUES
    (?, ?, ?, ?)
`

type CreateDocumentParams struct {
	PlayerID  int64
	Name      string
	Url       string
	ExpiresAt time.Time
}

func (q *Queries) CreateDocument(ctx context.Context, arg CreateDocumentParams) error {
	_, err := q.db.ExecContext(ctx, createDocument,
		arg.PlayerID,
		arg.Name,
		arg.Url,
		arg.ExpiresAt,
	)
	return err
}

const getDocument = `-- name: GetDocument :one
SELECT
    id, player_id, name, url, expires_at, created_at
FROM
    documents
WHERE
    id = ?
`

func (q *Queries) GetDocument(ctx context.Context, id int64) (Document, error) {
	row := q.db.QueryRowContext(ctx, getDocument, id)
	var i Document
	err := row.Scan(
		&i.ID,
		&i.PlayerID,
		&i.Name,
		&i.Url,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const listDocuments = `-- name: ListDocuments :many
SELECT
    id, player_id, name, url, expires_at, created_at
FROM
    documents
WHERE
    player_id = ?
`

func (q *Queries) ListDocuments(ctx context.Context, playerID int64) ([]Document, error) {
	rows, err := q.db.QueryContext(ctx, listDocuments, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Document
	for rows.Next() {
		var i Document
		if err := rows.Scan(
			&i.ID,
			&i.PlayerID,
			&i.Name,
			&i.Url,
			&i.ExpiresAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
