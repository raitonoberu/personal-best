// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: registration.sql

package sqlc

import (
	"context"
)

const createRegistration = `-- name: CreateRegistration :exec
INSERT INTO
    registrations (competition_id, player_id, is_approved, is_dropped)
VALUES
    (?, ?, ?, ?)
`

type CreateRegistrationParams struct {
	CompetitionID int64
	PlayerID      int64
	IsApproved    bool
	IsDropped     bool
}

func (q *Queries) CreateRegistration(ctx context.Context, arg CreateRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, createRegistration,
		arg.CompetitionID,
		arg.PlayerID,
		arg.IsApproved,
		arg.IsDropped,
	)
	return err
}

const deleteRegistration = `-- name: DeleteRegistration :exec
DELETE FROM
    registrations
WHERE
    competition_id = ? AND player_id = ?
`

type DeleteRegistrationParams struct {
	CompetitionID int64
	PlayerID      int64
}

func (q *Queries) DeleteRegistration(ctx context.Context, arg DeleteRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, deleteRegistration, arg.CompetitionID, arg.PlayerID)
	return err
}

const getRegistration = `-- name: GetRegistration :one
SELECT
    competition_id, player_id, is_approved, is_dropped, created_at
FROM
    registrations
WHERE
    player_id = ? AND competition_id = ?
`

type GetRegistrationParams struct {
	PlayerID      int64
	CompetitionID int64
}

func (q *Queries) GetRegistration(ctx context.Context, arg GetRegistrationParams) (Registration, error) {
	row := q.db.QueryRowContext(ctx, getRegistration, arg.PlayerID, arg.CompetitionID)
	var i Registration
	err := row.Scan(
		&i.CompetitionID,
		&i.PlayerID,
		&i.IsApproved,
		&i.IsDropped,
		&i.CreatedAt,
	)
	return i, err
}

const listApprovedCompetitionPlayers = `-- name: ListApprovedCompetitionPlayers :many
SELECT
    users.id, users.role_id, users.email, users.password, users.first_name, users.last_name, users.middle_name, users.created_at,
    players.user_id, players.birth_date, players.is_male, players.phone, players.telegram, players.preparation, players.position
FROM
    registrations
JOIN
    users ON (users.id = registrations.player_id)
JOIN
    players ON (players.user_id = registrations.player_id)
WHERE
    competition_id = ? AND is_approved = TRUE
`

type ListApprovedCompetitionPlayersRow struct {
	User   User
	Player Player
}

func (q *Queries) ListApprovedCompetitionPlayers(ctx context.Context, competitionID int64) ([]ListApprovedCompetitionPlayersRow, error) {
	rows, err := q.db.QueryContext(ctx, listApprovedCompetitionPlayers, competitionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListApprovedCompetitionPlayersRow
	for rows.Next() {
		var i ListApprovedCompetitionPlayersRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.RoleID,
			&i.User.Email,
			&i.User.Password,
			&i.User.FirstName,
			&i.User.LastName,
			&i.User.MiddleName,
			&i.User.CreatedAt,
			&i.Player.UserID,
			&i.Player.BirthDate,
			&i.Player.IsMale,
			&i.Player.Phone,
			&i.Player.Telegram,
			&i.Player.Preparation,
			&i.Player.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCompetitionPlayers = `-- name: ListCompetitionPlayers :many
SELECT
    users.id, users.role_id, users.email, users.password, users.first_name, users.last_name, users.middle_name, users.created_at
FROM
    registrations
JOIN
    users ON (users.id = registrations.player_id)
WHERE
    competition_id = ? AND is_approved = TRUE AND is_dropped = FALSE
`

type ListCompetitionPlayersRow struct {
	User User
}

func (q *Queries) ListCompetitionPlayers(ctx context.Context, competitionID int64) ([]ListCompetitionPlayersRow, error) {
	rows, err := q.db.QueryContext(ctx, listCompetitionPlayers, competitionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCompetitionPlayersRow
	for rows.Next() {
		var i ListCompetitionPlayersRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.RoleID,
			&i.User.Email,
			&i.User.Password,
			&i.User.FirstName,
			&i.User.LastName,
			&i.User.MiddleName,
			&i.User.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCompetitionRegistrations = `-- name: ListCompetitionRegistrations :many
SELECT
    registrations.competition_id, registrations.player_id, registrations.is_approved, registrations.is_dropped, registrations.created_at, users.id, users.role_id, users.email, users.password, users.first_name, users.last_name, users.middle_name, users.created_at, players.user_id, players.birth_date, players.is_male, players.phone, players.telegram, players.preparation, players.position
FROM
    registrations
    JOIN users ON (users.id = registrations.player_id)
    JOIN players ON (players.user_id = registrations.player_id)
WHERE
    competition_id = ?
`

type ListCompetitionRegistrationsRow struct {
	Registration Registration
	User         User
	Player       Player
}

func (q *Queries) ListCompetitionRegistrations(ctx context.Context, competitionID int64) ([]ListCompetitionRegistrationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listCompetitionRegistrations, competitionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCompetitionRegistrationsRow
	for rows.Next() {
		var i ListCompetitionRegistrationsRow
		if err := rows.Scan(
			&i.Registration.CompetitionID,
			&i.Registration.PlayerID,
			&i.Registration.IsApproved,
			&i.Registration.IsDropped,
			&i.Registration.CreatedAt,
			&i.User.ID,
			&i.User.RoleID,
			&i.User.Email,
			&i.User.Password,
			&i.User.FirstName,
			&i.User.LastName,
			&i.User.MiddleName,
			&i.User.CreatedAt,
			&i.Player.UserID,
			&i.Player.BirthDate,
			&i.Player.IsMale,
			&i.Player.Phone,
			&i.Player.Telegram,
			&i.Player.Preparation,
			&i.Player.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlayerRegistrations = `-- name: ListPlayerRegistrations :many
SELECT
    users.id, users.role_id, users.email, users.password, users.first_name, users.last_name, users.middle_name, users.created_at,
    competitions.id, competitions.trainer_id, competitions.name, competitions.description, competitions.tours, competitions.age, competitions.size, competitions.closes_at, competitions.created_at,
    registrations.competition_id, registrations.player_id, registrations.is_approved, registrations.is_dropped, registrations.created_at,
    COUNT() OVER() as total
FROM
    registrations
    JOIN competitions ON competitions.id = registrations.competition_id
    JOIN users ON users.id = competitions.trainer_id
WHERE
    registrations.player_id = ?
LIMIT
    ? OFFSET ?
`

type ListPlayerRegistrationsParams struct {
	PlayerID int64
	Limit    int64
	Offset   int64
}

type ListPlayerRegistrationsRow struct {
	User         User
	Competition  Competition
	Registration Registration
	Total        int64
}

func (q *Queries) ListPlayerRegistrations(ctx context.Context, arg ListPlayerRegistrationsParams) ([]ListPlayerRegistrationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPlayerRegistrations, arg.PlayerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPlayerRegistrationsRow
	for rows.Next() {
		var i ListPlayerRegistrationsRow
		if err := rows.Scan(
			&i.User.ID,
			&i.User.RoleID,
			&i.User.Email,
			&i.User.Password,
			&i.User.FirstName,
			&i.User.LastName,
			&i.User.MiddleName,
			&i.User.CreatedAt,
			&i.Competition.ID,
			&i.Competition.TrainerID,
			&i.Competition.Name,
			&i.Competition.Description,
			&i.Competition.Tours,
			&i.Competition.Age,
			&i.Competition.Size,
			&i.Competition.ClosesAt,
			&i.Competition.CreatedAt,
			&i.Registration.CompetitionID,
			&i.Registration.PlayerID,
			&i.Registration.IsApproved,
			&i.Registration.IsDropped,
			&i.Registration.CreatedAt,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRegistration = `-- name: UpdateRegistration :exec
UPDATE
    registrations
SET
    is_approved = coalesce(?1, is_approved),
    is_dropped = coalesce(?2, is_dropped)
WHERE
    player_id = ?3 AND competition_id = ?4
`

type UpdateRegistrationParams struct {
	IsApproved    *bool
	IsDropped     *bool
	PlayerID      int64
	CompetitionID int64
}

func (q *Queries) UpdateRegistration(ctx context.Context, arg UpdateRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, updateRegistration,
		arg.IsApproved,
		arg.IsDropped,
		arg.PlayerID,
		arg.CompetitionID,
	)
	return err
}
