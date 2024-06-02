// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tickets.sql

package db

import (
	"context"
	"database/sql"

	"github.com/guregu/null"
)

const countTicketByStatusID = `-- name: CountTicketByStatusID :one
SELECT
  COUNT(*)
FROM
  tickets
WHERE
  status_id = ?
`

func (q *Queries) CountTicketByStatusID(ctx context.Context, statusID uint32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTicketByStatusID, statusID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTicket = `-- name: CreateTicket :exec
INSERT INTO
  tickets (
    status_id,
    title,
    description,
    contact,
    sort_order,
    created_at
  )
VALUES
  (?, ?, ?, ?, ?, NOW())
`

type CreateTicketParams struct {
	StatusID    uint32         `db:"status_id" json:"status_id"`
	Title       null.String    `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	Contact     null.String    `db:"contact" json:"contact"`
	SortOrder   uint32         `db:"sort_order" json:"sort_order"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) error {
	_, err := q.db.ExecContext(ctx, createTicket,
		arg.StatusID,
		arg.Title,
		arg.Description,
		arg.Contact,
		arg.SortOrder,
	)
	return err
}

const getLastInsertTicketByStatusID = `-- name: GetLastInsertTicketByStatusID :one
SELECT
  id, status_id, title, description, contact, sort_order, created_at, updated_at
FROM
  tickets
WHERE
  tickets.status_id = ?
  AND id = (
    SELECT
      LAST_INSERT_ID()
    FROM
      tickets AS t
    LIMIT
      1
  )
`

func (q *Queries) GetLastInsertTicketByStatusID(ctx context.Context, statusID uint32) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, getLastInsertTicketByStatusID, statusID)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.StatusID,
		&i.Title,
		&i.Description,
		&i.Contact,
		&i.SortOrder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTicketsByStatusID = `-- name: GetTicketsByStatusID :many
SELECT
  id, status_id, title, description, contact, sort_order, created_at, updated_at
FROM
  tickets
WHERE
  status_id = ?
`

func (q *Queries) GetTicketsByStatusID(ctx context.Context, statusID uint32) ([]Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByStatusID, statusID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ticket{}
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.StatusID,
			&i.Title,
			&i.Description,
			&i.Contact,
			&i.SortOrder,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateTicket = `-- name: UpdateTicket :exec
UPDATE
  tickets
SET
  title = ?,
  description = ?,
  contact = ?,
  updated_at = NOW()
WHERE
  id = ?
`

type UpdateTicketParams struct {
	Title       null.String    `db:"title" json:"title"`
	Description sql.NullString `db:"description" json:"description"`
	Contact     null.String    `db:"contact" json:"contact"`
	ID          uint64         `db:"id" json:"id"`
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) error {
	_, err := q.db.ExecContext(ctx, updateTicket,
		arg.Title,
		arg.Description,
		arg.Contact,
		arg.ID,
	)
	return err
}
