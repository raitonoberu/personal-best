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

-- name: GetMatchPlayersToUpdateScore :many
SELECT
    *
FROM
    match_players
JOIN
    matches ON match_id = matches.id
WHERE
    player_id = ? AND matches.competition_id = ?
ORDER BY
    match_id DESC
LIMIT 2;

-- name: SetMatchPlayerWinLoseScores :exec
UPDATE
    match_players
SET
    win_score = ?, lose_score = ?
WHERE
    match_id = ? AND player_id = ?;

-- name: GetMatchPlayerLastScores :one
SELECT
    *
FROM
    match_players
JOIN
    matches ON match_id = matches.id
WHERE
    player_id = ? AND win_score IS NOT NULL AND matches.competition_id = ?
ORDER BY
    match_id DESC
LIMIT 1;
