// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: tfe_run.sql

package db

import (
	"context"
)

const createTFERun = `-- name: CreateTFERun :one
INSERT INTO tfe_run (id, configuration_version_id, status)
VALUES ($1, $2, $3)
RETURNING id, configuration_version_id, status, created_at, updated_at
`

type CreateTFERunParams struct {
	ID                     string
	ConfigurationVersionID string
	Status                 TFERunStatus
}

func (q *Queries) CreateTFERun(ctx context.Context, arg CreateTFERunParams) (TFERun, error) {
	row := q.db.QueryRow(ctx, createTFERun, arg.ID, arg.ConfigurationVersionID, arg.Status)
	var i TFERun
	err := row.Scan(
		&i.ID,
		&i.ConfigurationVersionID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTFERun = `-- name: DeleteTFERun :exec
DELETE FROM tfe_run
WHERE id = $1
`

func (q *Queries) DeleteTFERun(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteTFERun, id)
	return err
}

const getTFERun = `-- name: GetTFERun :one
SELECT id, configuration_version_id, status, created_at, updated_at
FROM tfe_run
WHERE id = $1
`

func (q *Queries) GetTFERun(ctx context.Context, id string) (TFERun, error) {
	row := q.db.QueryRow(ctx, getTFERun, id)
	var i TFERun
	err := row.Scan(
		&i.ID,
		&i.ConfigurationVersionID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTFERuns = `-- name: ListTFERuns :many
SELECT id, configuration_version_id, status, created_at, updated_at
FROM tfe_run
`

func (q *Queries) ListTFERuns(ctx context.Context) ([]TFERun, error) {
	rows, err := q.db.Query(ctx, listTFERuns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TFERun
	for rows.Next() {
		var i TFERun
		if err := rows.Scan(
			&i.ID,
			&i.ConfigurationVersionID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTFERun = `-- name: UpdateTFERun :one
UPDATE tfe_run
SET configuration_version_id = $2,
  status = $3,
  updated_at = NOW()
WHERE id = $1
RETURNING id, configuration_version_id, status, created_at, updated_at
`

type UpdateTFERunParams struct {
	ID                     string
	ConfigurationVersionID string
	Status                 TFERunStatus
}

func (q *Queries) UpdateTFERun(ctx context.Context, arg UpdateTFERunParams) (TFERun, error) {
	row := q.db.QueryRow(ctx, updateTFERun, arg.ID, arg.ConfigurationVersionID, arg.Status)
	var i TFERun
	err := row.Scan(
		&i.ID,
		&i.ConfigurationVersionID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTFERunStatus = `-- name: UpdateTFERunStatus :one
UPDATE tfe_run
SET status = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING id, configuration_version_id, status, created_at, updated_at
`

type UpdateTFERunStatusParams struct {
	ID     string
	Status TFERunStatus
}

func (q *Queries) UpdateTFERunStatus(ctx context.Context, arg UpdateTFERunStatusParams) (TFERun, error) {
	row := q.db.QueryRow(ctx, updateTFERunStatus, arg.ID, arg.Status)
	var i TFERun
	err := row.Scan(
		&i.ID,
		&i.ConfigurationVersionID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}