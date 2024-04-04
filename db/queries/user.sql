-- name: CreateUser :one
INSERT INTO
    users (name, email, password, is_trainer, birth_date)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?
LIMIT
    1;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT
    sqlc.embed(users),
    COUNT() OVER() as total
FROM
    users
LIMIT
    ? OFFSET ?;

-- name: UpdateUser :one
UPDATE
    users
SET
    name = ?,
    email = ?,
    password = ?
WHERE
    id = ? RETURNING *;

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?;