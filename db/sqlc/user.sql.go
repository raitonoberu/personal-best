// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package sqlc

import (
	"context"
	"time"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO
    players (user_id, birth_date, is_male, phone, telegram)
VALUES
    (?, ?, ?, ?, ?)
RETURNING user_id, birth_date, is_male, phone, telegram, is_verified, preparation, position
`

type CreatePlayerParams struct {
	UserID    int64
	BirthDate time.Time
	IsMale    bool
	Phone     string
	Telegram  string
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, createPlayer,
		arg.UserID,
		arg.BirthDate,
		arg.IsMale,
		arg.Phone,
		arg.Telegram,
	)
	var i Player
	err := row.Scan(
		&i.UserID,
		&i.BirthDate,
		&i.IsMale,
		&i.Phone,
		&i.Telegram,
		&i.IsVerified,
		&i.Preparation,
		&i.Position,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (role_id, email, password, first_name, last_name, middle_name)
VALUES
    (?, ?, ?, ?, ?, ?)
RETURNING id, role_id, email, password, first_name, last_name, middle_name, created_at
`

type CreateUserParams struct {
	RoleID     int64
	Email      string
	Password   string
	FirstName  string
	LastName   string
	MiddleName string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.RoleID,
		arg.Email,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT
    id, role_id, email, password, first_name, last_name, middle_name, created_at
FROM
    users
WHERE
    id = ?
LIMIT
    1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id, role_id, email, password, first_name, last_name, middle_name, created_at
FROM
    users
WHERE
    email = ?
LIMIT
    1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.RoleID,
		&i.Email,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE
    users
SET
    first_name = coalesce(?1, first_name),
    last_name = coalesce(?2, last_name),
    middle_name = coalesce(?3, middle_name),
    email = coalesce(?4, email),
    password = coalesce(?5, password)
WHERE
    id = ?6
`

type UpdateUserParams struct {
	FirstName  *string
	LastName   *string
	MiddleName *string
	Email      *string
	Password   *string
	ID         int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	return err
}
