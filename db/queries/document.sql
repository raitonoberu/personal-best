-- name: CreateDocument :exec
INSERT INTO
    documents (player_id, name, url, expires_at)
VALUES
    (?, ?, ?, ?);

-- name: GetDocument :one
SELECT
    *
FROM
    documents
WHERE
    id = ?;

-- name: ListDocuments :many
SELECT
    *
FROM
    documents
WHERE
    player_id = ?;

-- name: DeleteDocument :exec
DELETE FROM
    documents
WHERE
    id = ?;
