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
    sqlc.embed(users), sqlc.embed(user_players)
FROM
    users
LEFT JOIN
    user_players ON users.id = user_players.user_id
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
    role_id = coalesce(sqlc.narg('role_id'), role_id),
    email = coalesce(sqlc.narg('email'), email),
    password = coalesce(sqlc.narg('password'), password),
    first_name = coalesce(sqlc.narg('first_name'), first_name),
    last_name = coalesce(sqlc.narg('last_name'), last_name),
    middle_name = coalesce(sqlc.narg('middle_name'), middle_name)
WHERE
    id = sqlc.arg('id');

-- name: UpdatePlayer :exec
UPDATE
    players
SET
    birth_date = coalesce(sqlc.narg('birth_date'), birth_date),
    is_male = coalesce(sqlc.narg('is_male'), is_male),
    phone = coalesce(sqlc.narg('phone'), phone),
    telegram = coalesce(sqlc.narg('telegram'), telegram)
WHERE
    user_id = sqlc.arg('user_id');

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?;

