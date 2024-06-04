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
    users.id, users.role_id, users.email, users.password, users.first_name, users.last_name, users.middle_name, users.created_at, registrations.competition_id, registrations.player_id, registrations.is_approved, registrations.is_dropped, registrations.created_at
FROM
    registrations
JOIN
    users ON (users.id = registrations.player_id)
WHERE
    competition_id = ?
`

type ListCompetitionRegistrationsRow struct {
	User         User
	Registration Registration
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
			&i.User.ID,
			&i.User.RoleID,
			&i.User.Email,
			&i.User.Password,
			&i.User.FirstName,
			&i.User.LastName,
			&i.User.MiddleName,
			&i.User.CreatedAt,
			&i.Registration.CompetitionID,
			&i.Registration.PlayerID,
			&i.Registration.IsApproved,
			&i.Registration.IsDropped,
			&i.Registration.CreatedAt,
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
    competitions.id, competitions.trainer_id, competitions.name, competitions.description, competitions.tours, competitions.age, competitions.size, competitions.closes_at, competitions.created_at, registrations.competition_id, registrations.player_id, registrations.is_approved, registrations.is_dropped, registrations.created_at
FROM
    registrations
JOIN
    competitions ON (competitions.id = registrations.competitions_id)
WHERE
    player_id = ?
`

type ListPlayerRegistrationsRow struct {
	Competition  Competition
	Registration Registration
}

func (q *Queries) ListPlayerRegistrations(ctx context.Context, playerID int64) ([]ListPlayerRegistrationsRow, error) {
	rows, err := q.db.QueryContext(ctx, listPlayerRegistrations, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPlayerRegistrationsRow
	for rows.Next() {
		var i ListPlayerRegistrationsRow
		if err := rows.Scan(
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
