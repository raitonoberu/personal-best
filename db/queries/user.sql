-- name: CreateUser :one
INSERT INTO
    users (role_id, email, password, first_name, last_name, middle_name)
VALUES
    (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: CreatePlayer :one
INSERT INTO
    players (user_id, birth_date, is_male, phone, telegram)
VALUES
    (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUser :one
SELECT
    sqlc.embed(users), sqlc.embed(players), sqlc.embed(roles)
FROM
    users
LEFT JOIN
    players ON users.id = players.user_id
JOIN
    roles ON users.role_id = roles.id
WHERE
    users.id = ?
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

-- name: UpdateUser :exec
UPDATE
    users
SET
    first_name = coalesce(sqlc.narg('first_name'), first_name),
    last_name = coalesce(sqlc.narg('last_name'), last_name),
    middle_name = coalesce(sqlc.narg('middle_name'), middle_name),
    email = coalesce(sqlc.narg('email'), email),
    password = coalesce(sqlc.narg('password'), password)
WHERE
    id = sqlc.arg('id');

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?;

