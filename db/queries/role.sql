-- name: GetRole :one
SELECT
    *
FROM
    roles
WHERE
    id = ?;

-- name: GetUserRole :one
SELECT
    roles.*
FROM
    users
JOIN
    roles ON users.role_id = roles.id
WHERE
    users.id = ?;

-- name: ListRoles :many
SELECT
    *
FROM
    roles;

