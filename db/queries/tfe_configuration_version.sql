-- name: GetTFEConfigurationVersion :one
SELECT *
FROM tfe_configuration_version
WHERE id = $1;

-- name: CreateTFEConfigurationVersion :one
INSERT INTO tfe_configuration_version (
    id,
    auto_queue_runs,
    speculative,
    provisional,
    STATUS,
    upload_data
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTFEConfigurationVersion :one
UPDATE tfe_configuration_version
SET auto_queue_runs = $2,
  speculative = $3,
  provisional = $4,
  STATUS = $5,
  upload_data = $6,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTFEConfigurationVersion :exec
DELETE FROM tfe_configuration_version
WHERE id = $1;

-- name: UploadTFEConfigurationVersion :one
UPDATE tfe_configuration_version
SET STATUS = 'uploaded',
  upload_data = $2
WHERE id = $1
RETURNING *;