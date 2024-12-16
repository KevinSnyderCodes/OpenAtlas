CREATE TYPE tfe_configuration_version_status AS ENUM ('pending', 'uploaded');

CREATE TABLE tfe_configuration_version (
  -- Keys
  id TEXT NOT NULL PRIMARY KEY,
  -- Columns
  auto_queue_runs BOOLEAN NOT NULL,
  speculative BOOLEAN NOT NULL,
  provisional BOOLEAN NOT NULL,
  status tfe_configuration_version_status NOT NULL,
  upload_data BYTEA,
  -- Timestamps
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TYPE tfe_run_status AS ENUM (
  -- TODO: Order the same was as in the documentation
  -- https://developer.hashicorp.com/terraform/enterprise/api-docs/run#run-states
  'applied',
  'applying',
  'apply_queued',
  'canceled',
  'confirmed',
  'cost_estimated',
  'cost_estimating',
  'discarded',
  'errored',
  'fetching',
  'fetching_completed',
  'pending',
  'planned',
  'planned_and_finished',
  'planning',
  'plan_queued',
  'policy_checked',
  'policy_checking',
  'policy_override',
  'policy_soft_failed',
  'post_plan_awaiting_decision',
  'post_plan_completed',
  'post_plan_running',
  'pre_plan_completed',
  'pre_plan_running',
  'queuing'
);

CREATE TABLE tfe_run (
  -- Keys
  id TEXT NOT NULL PRIMARY KEY,
  configuration_version_id TEXT NOT NULL REFERENCES tfe_configuration_version(id),
  -- Columns
  status tfe_run_status NOT NULL,
  -- Timestamps
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TYPE tfe_plan_status AS ENUM (
  'pending',
  'managed_queued',
  'queued',
  'running',
  'errored',
  'canceled',
  'finished',
  'unreachable'
);

CREATE TABLE tfe_plan (
  -- Keys
  id TEXT NOT NULL PRIMARY KEY,
  run_id TEXT NOT NULL REFERENCES tfe_run(id),
  -- Columns
  status tfe_plan_status NOT NULL,
  log_read_url TEXT NOT NULL,
  -- Timestamps
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE
);