-- name: CreateRegistration :exec
INSERT INTO
    registrations (competition_id, player_id, is_approved, is_dropped)
VALUES
    (?, ?, ?, ?);

-- name: DeleteRegistration :exec
DELETE FROM
    registrations
WHERE
    competition_id = ? AND player_id = ?;

-- name: ListCompetitionRegistrations :many
SELECT
    sqlc.embed(users), sqlc.embed(registrations)
FROM
    registrations
JOIN
    users ON (users.id = registrations.player_id)
WHERE
    competition_id = ?;

-- name: ListPlayerRegistrations :many
SELECT
    sqlc.embed(competitions), sqlc.embed(registrations)
FROM
    registrations
JOIN
    competitions ON (competitions.id = registrations.competitions_id)
WHERE
    player_id = ?;

-- name: ListCompetitionPlayers :many
SELECT
    sqlc.embed(users)
FROM
    registrations
JOIN
    users ON (users.id = registrations.player_id)
WHERE
    competition_id = ? AND is_approved = TRUE AND is_dropped = FALSE;

-- name: UpdateRegistration :exec
UPDATE
    registrations
SET
    is_approved = coalesce(sqlc.narg('is_approved'), is_approved),
    is_dropped = coalesce(sqlc.narg('is_dropped'), is_dropped)
WHERE
    player_id = sqlc.arg('player_id') AND competition_id = sqlc.arg('competition_id');
