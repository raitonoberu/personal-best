-- name: CreateCompetition :one
INSERT INTO
    competitions (trainer_id, name, description, start_date, tours, age, size, closes_at)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: GetCompetition :one
SELECT
    sqlc.embed(users),
    sqlc.embed(competitions),
    sqlc.embed(competition_days)
FROM
    competitions
    JOIN users ON users.id = competitions.trainer_id
    JOIN competition_days ON competition_id = competitions.id
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

-- name: UpdateCompetition :exec
UPDATE
    competitions
SET
    name = coalesce(sqlc.narg('name'), name),
    description = coalesce(sqlc.narg('description'), description),
    closes_at = coalesce(sqlc.narg('closes_at'), closes_at)
    -- more?
WHERE
    id = sqlc.arg('id');

-- name: DeleteCompetition :exec
DELETE FROM
    competitions
WHERE
    id = ?;
