-- name: GetTFERun :one
SELECT *
FROM tfe_run
WHERE id = $1;

-- name: ListTFERuns :many
SELECT *
FROM tfe_run;

-- name: CreateTFERun :one
INSERT INTO tfe_run (id, configuration_version_id, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateTFERun :one
UPDATE tfe_run
SET configuration_version_id = $2,
  status = $3,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateTFERunStatus :one
UPDATE tfe_run
SET status = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTFERun :exec
DELETE FROM tfe_run
WHERE id = $1;