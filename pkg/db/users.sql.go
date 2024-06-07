// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"

	null "github.com/guregu/null/v5"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO
  users (name, lastname, email, password, created_at)
VALUES
  (?, ?, ?, ?, NOW())
`

type CreateUserParams struct {
	Name     null.String `db:"name" json:"name"`
	Lastname null.String `db:"lastname" json:"lastname"`
	Email    null.String `db:"email" json:"email"`
	Password null.String `db:"password" json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Name,
		arg.Lastname,
		arg.Email,
		arg.Password,
	)
	return err
}

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT
  id, name, lastname, email, password, created_at, updated_at
FROM
  users
WHERE
  email = ?
LIMIT
  1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email null.String) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT
  id, name, lastname, email, password, created_at, updated_at
FROM
  users
WHERE
  id = ?
`

func (q *Queries) FindUserByID(ctx context.Context, id uint64) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLastInsertUser = `-- name: GetLastInsertUser :one
SELECT
  id, name, lastname, email, password, created_at, updated_at
FROM
  users
WHERE
  id = (
    SELECT
      LAST_INSERT_ID()
    FROM
      users AS u
    LIMIT
      1
  )
`

func (q *Queries) GetLastInsertUser(ctx context.Context) (User, error) {
	row := q.db.QueryRowContext(ctx, getLastInsertUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLastInsertUserID = `-- name: GetLastInsertUserID :one
SELECT
  LAST_INSERT_ID()
FROM
  users
LIMIT
  1
`

func (q *Queries) GetLastInsertUserID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getLastInsertUserID)
	var last_insert_id int64
	err := row.Scan(&last_insert_id)
	return last_insert_id, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
  id, name, lastname, email, password, created_at, updated_at
FROM
  users
WHERE
  id = ?
LIMIT
  1
`

func (q *Queries) GetUserByID(ctx context.Context, id uint64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE
  users
SET
  name = ?,
  lastname = ?,
  email = ?,
  password = ?,
  updated_at = NOW()
WHERE
  id = ?
`

type UpdateUserParams struct {
	Name     null.String `db:"name" json:"name"`
	Lastname null.String `db:"lastname" json:"lastname"`
	Email    null.String `db:"email" json:"email"`
	Password null.String `db:"password" json:"password"`
	ID       uint64      `db:"id" json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Lastname,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	return err
}
