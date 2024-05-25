-- name: CreateCompetition :one
INSERT INTO
    competitions (trainer_id, name, description, tours, age, size, closes_at)
VALUES
    (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: CreateCompetitionDay :one
INSERT INTO
    competition_days (competition_id, date, start_time, end_time)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetCompetition :one
SELECT
    sqlc.embed(competitions),
    sqlc.embed(users)
FROM
    competitions
    JOIN users ON users.id = competitions.trainer_id
WHERE
    competitions.id = ?
LIMIT
    1;

-- name: GetCompetitionDays :many
SELECT
    *
FROM
    competition_days
WHERE
    competition_id = ?
ORDER BY date;

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

