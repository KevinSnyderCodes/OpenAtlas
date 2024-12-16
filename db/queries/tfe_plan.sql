-- name: GetTFEPlan :one
SELECT *
FROM tfe_plan
WHERE id = $1;

-- name: GetTFEPlanByRunID :one
SELECT *
FROM tfe_plan
WHERE run_id = $1;

-- name: ListTFEPlans :many
SELECT *
FROM tfe_plan;

-- name: CreateTFEPlan :one
INSERT INTO tfe_plan (id, run_id, status, log_read_url)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateTFEPlan :one
UPDATE tfe_plan
SET log_read_url = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateTFEPlanStatus :one
UPDATE tfe_plan
SET status = $2,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTFEPlan :exec
DELETE FROM tfe_plan
WHERE id = $1;