CREATE TYPE tfe_configuration_version_status AS ENUM ('pending', 'uploaded');

CREATE TABLE tfe_configuration_version (
  id TEXT NOT NULL PRIMARY KEY,
  auto_queue_runs BOOLEAN NOT NULL,
  speculative BOOLEAN NOT NULL,
  provisional BOOLEAN NOT NULL,
  status tfe_configuration_version_status NOT NULL,
  upload_data BYTEA
);

CREATE TYPE tfe_run_status AS ENUM (
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
  id TEXT NOT NULL PRIMARY KEY,
  configuration_version_id TEXT NOT NULL REFERENCES tfe_configuration_version(id),
  status tfe_run_status NOT NULL
);