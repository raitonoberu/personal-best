-- name: CreateCompetition :one
INSERT INTO
    competitions (name, description, start_date, trainer_id)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetCompetition :one
SELECT
    sqlc.embed(users),
    sqlc.embed(competitions)
FROM
    competitions
    JOIN users ON users.id = competitions.trainer_id
WHERE
    competitions.id = ?
LIMIT
    1;

-- name: ListCompetitions :many
SELECT
    sqlc.embed(users),
    sqlc.embed(competitions),
    COUNT() OVER() as total
FROM
    competitions
    JOIN users ON users.id = competitions.trainer_id
LIMIT
    ? OFFSET ?;

-- name: ListCompetitionsByTrainer :many
SELECT
    *
FROM
    competitions
WHERE
    trainer_id = ?
LIMIT
    ? OFFSET ?;

-- name: UpdateCompetition :one
UPDATE
    competitions
SET
    name = ?,
    description = ?,
    start_date = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteCompetition :exec
DELETE FROM
    competitions
WHERE
    id = ?;