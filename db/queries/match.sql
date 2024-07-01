-- name: GetMatch :one
SELECT
    *
FROM
    matches
WHERE
    id = ?;

-- name: GetLastMatch :one
SELECT
    *
FROM
    matches
WHERE
    competition_id = ? AND left_score IS NOT NULL AND right_score IS NOT NULL
ORDER BY
    start_time DESC
LIMIT 1;

-- name: GetNextMatch :one
SELECT
    *
FROM
    matches
WHERE
    competition_id = ? AND left_score IS NULL AND right_score IS NULL
ORDER BY
    start_time
LIMIT 1;

-- name: CreateMatch :execlastid
INSERT INTO
    matches (competition_id, start_time)
VALUES
    (?, ?);

-- name: ListMatches :many
SELECT
    sqlc.embed(matches),
    COUNT() OVER() as total
FROM
    matches
WHERE
    competition_id = ?
ORDER BY
    start_time
LIMIT
    ? OFFSET ?;

-- name: ListAllMatches :many
SELECT
    *
FROM
    matches
WHERE
    competition_id = ?
ORDER BY
    start_time;

-- name: DeleteMatches :exec
DELETE FROM
    matches
WHERE
    competition_id = ?;

-- name: UpdateMatchScore :exec
UPDATE
    matches
SET
    left_score = ?, right_score = ?
WHERE
    id = ?;

