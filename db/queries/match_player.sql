-- name: ListMatchPlayers :many
SELECT
    *
FROM
    match_players
WHERE
    match_id = ?;

-- name: ListMatchPlayersBatch :many
SELECT
    *
FROM
    match_players
WHERE
    match_id IN (sqlc.slice(match_ids));

-- name: ListMatchPlayersWithPlayersBatch :many
SELECT
    sqlc.embed(match_players),
    sqlc.embed(users),
    sqlc.embed(players)
FROM
    match_players
JOIN
    users ON match_players.player_id == users.id
JOIN
    players ON match_players.player_id == players.user_id
WHERE
    match_id IN (sqlc.slice(match_ids));

-- name: CreateMatchPlayer :exec
INSERT INTO
    match_players (match_id, player_id, position, team)
VALUES
    (?, ?, ?, ?);

-- name: DeleteMatchPlayers :exec
DELETE FROM
    match_players
WHERE
    match_id = ?;


