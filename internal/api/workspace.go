package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

var _ Workspaces = (*DefaultWorkspaces)(nil)

type Workspaces interface {
	Read(ctx context.Context, organization string, workspace string) (*WorkspaceReadResponse, error)
	Create(ctx context.Context, req *WorkspaceCreateRequest) (*WorkspaceCreateResponse, error)
}

type WorkspaceReadResponse jsonapi.Document[
	*jsonapi.Resource[*WorkspaceReadResponseResourceAttributes],
	WorkspaceReadResponseResourceAttributes,
]

type WorkspaceReadResponseResourceAttributes struct {
	Actions                     *WorkspaceReadResponseResourceAttributesActions           `json:"actions"`
	AllowDestroyPlan            bool                                                      `json:"allow-destroy-plan"`
	ApplyDurationAverage        int                                                       `json:"apply-duration-average"`
	AutoApply                   bool                                                      `json:"auto-apply"`
	AutoApplyRunTrigger         bool                                                      `json:"auto-apply-run-trigger"`
	AutoDestroyAt               any                                                       `json:"auto-destroy-at"`
	AutoDestroyActivityDuration any                                                       `json:"auto-destroy-activity-duration"`
	CreatedAt                   time.Time                                                 `json:"created-at"`
	Description                 string                                                    `json:"description"`
	Environment                 string                                                    `json:"environment"`
	ExecutionMode               string                                                    `json:"execution-mode"`
	FileTriggersEnabled         bool                                                      `json:"file-triggers-enabled"`
	GlobalRemoteState           bool                                                      `json:"global-remote-state"`
	LatestChangeAt              time.Time                                                 `json:"latest-change-at"`
	Locked                      bool                                                      `json:"locked"`
	Name                        string                                                    `json:"name"`
	Operations                  bool                                                      `json:"operations"`
	Permissions                 *WorkspaceReadResponseResourceAttributesPermissions       `json:"permissions"`
	PlanDurationAverage         int                                                       `json:"plan-duration-average"`
	PolicyCheckFailures         any                                                       `json:"policy-check-failures"`
	QueueAllRuns                bool                                                      `json:"queue-all-runs"`
	ResourceCount               int                                                       `json:"resource-count"`
	RunFailures                 int                                                       `json:"run-failures"`
	Source                      string                                                    `json:"source"`
	SourceName                  any                                                       `json:"source-name"`
	SourceURL                   any                                                       `json:"source-url"`
	SpeculativeEnabled          bool                                                      `json:"speculative-enabled"`
	StructuredRunOutputEnabled  bool                                                      `json:"structured-run-output-enabled"`
	TerraformVersion            string                                                    `json:"terraform-version"`
	TriggerPrefixes             []string                                                  `json:"trigger-prefixes"`
	UpdatedAt                   time.Time                                                 `json:"updated-at"`
	VCSRepo                     any                                                       `json:"vcs-repo"`
	VCSRepoIdentifier           any                                                       `json:"vcs-repo-identifier"`
	WorkingDirectory            any                                                       `json:"working-directory"`
	WorkspaceKPIsRunsCount      int                                                       `json:"workspace-kpis-runs-count"`
	SettingOverwrites           *WorkspaceReadResponseResourceAttributesSettingOverwrites `json:"setting-overwrites"`
}

type WorkspaceReadResponseResourceAttributesActions struct {
	IsDestroyable bool `json:"is-destroyable"`
}

type WorkspaceReadResponseResourceAttributesPermissions struct {
	CanCreateStateVersions bool `json:"can-create-state-versions"`
	CanDestroy             bool `json:"can-destroy"`
	CanForceUnlock         bool `json:"can-force-unlock"`
	CanLock                bool `json:"can-lock"`
	CanManageRunTasks      bool `json:"can-manage-run-tasks"`
	CanManageTags          bool `json:"can-manage-tags"`
	CanQueueApply          bool `json:"can-queue-apply"`
	CanQueueDestroy        bool `json:"can-queue-destroy"`
	CanQueueRun            bool `json:"can-queue-run"`
	CanReadSettings        bool `json:"can-read-settings"`
	CanReadStateVersions   bool `json:"can-read-state-versions"`
	CanReadVariable        bool `json:"can-read-variable"`
	CanUnlock              bool `json:"can-unlock"`
	CanUpdate              bool `json:"can-update"`
	CanUpdateVariable      bool `json:"can-update-variable"`
	CanForceDelete         bool `json:"can-force-delete"`
}

type WorkspaceReadResponseResourceAttributesSettingOverwrites struct {
	ExecutionMode bool `json:"execution-mode"`
	AgentPool     bool `json:"agent-pool"`
}

func (o *WorkspaceReadResponse) MarshalJSON() ([]byte, error) {
	type Alias WorkspaceReadResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type WorkspaceCreateRequest jsonapi.Document[
	*jsonapi.Resource[*WorkspaceCreateRequestResourceAttributes],
	WorkspaceCreateRequestResourceAttributes,
]

type WorkspaceCreateRequestResourceAttributes struct {
	Name          string    `json:"name"`
	ResourceCount int       `json:"resource-count"`
	UpdatedAt     time.Time `json:"updated-at"`
}

func (o *WorkspaceCreateRequest) UnmarshalJSON(data []byte) error {
	type Alias WorkspaceCreateRequest

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	*o = WorkspaceCreateRequest(alias)

	return nil
}

type WorkspaceCreateResponse jsonapi.Document[
	*jsonapi.Resource[*WorkspaceCreateResponseResourceAttributes],
	WorkspaceCreateResponseResourceAttributes,
]

type WorkspaceCreateResponseResourceAttributes struct {
	Actions                     *WorkspaceCreateResponseResourceAttributesActions           `json:"actions"`
	AllowDestroyPlan            bool                                                        `json:"allow-destroy-plan"`
	ApplyDurationAverage        any                                                         `json:"apply-duration-average"`
	AutoApply                   bool                                                        `json:"auto-apply"`
	AutoApplyRunTrigger         bool                                                        `json:"auto-apply-run-trigger"`
	AutoDestroyAt               any                                                         `json:"auto-destroy-at"`
	AutoDestroyActivityDuration any                                                         `json:"auto-destroy-activity-duration"`
	CreatedAt                   time.Time                                                   `json:"created-at"`
	Description                 any                                                         `json:"description"`
	Environment                 string                                                      `json:"environment"`
	ExecutionMode               string                                                      `json:"execution-mode"`
	FileTriggersEnabled         bool                                                        `json:"file-triggers-enabled"`
	GlobalRemoteState           bool                                                        `json:"global-remote-state"`
	LatestChangeAt              time.Time                                                   `json:"latest-change-at"`
	LastAssessmentResultAt      any                                                         `json:"last-assessment-result-at"`
	Locked                      bool                                                        `json:"locked"`
	Name                        string                                                      `json:"name"`
	Operations                  bool                                                        `json:"operations"`
	Permissions                 *WorkspaceCreateResponseResourceAttributesPermissions       `json:"permissions"`
	PlanDurationAverage         any                                                         `json:"plan-duration-average"`
	PolicyCheckFailures         any                                                         `json:"policy-check-failures"`
	QueueAllRuns                bool                                                        `json:"queue-all-runs"`
	ResourceCount               int                                                         `json:"resource-count"`
	RunFailures                 any                                                         `json:"run-failures"`
	Source                      string                                                      `json:"source"`
	SourceName                  any                                                         `json:"source-name"`
	SourceURL                   any                                                         `json:"source-url"`
	SpeculativeEnabled          bool                                                        `json:"speculative-enabled"`
	StructuredRunOutputEnabled  bool                                                        `json:"structured-run-output-enabled"`
	TerraformVersion            string                                                      `json:"terraform-version"`
	TriggerPrefixes             []string                                                    `json:"trigger-prefixes"`
	UpdatedAt                   time.Time                                                   `json:"updated-at"`
	VCSRepo                     any                                                         `json:"vcs-repo"`
	VCSRepoIdentifier           any                                                         `json:"vcs-repo-identifier"`
	WorkingDirectory            any                                                         `json:"working-directory"`
	WorkspaceKPIsRunsCount      any                                                         `json:"workspace-kpis-runs-count"`
	AssessmentsEnabled          bool                                                        `json:"assessments-enabled"`
	SettingOverwrites           *WorkspaceCreateResponseResourceAttributesSettingOverwrites `json:"setting-overwrites"`
}

type WorkspaceCreateResponseResourceAttributesActions struct {
	IsDestroyable bool `json:"is-destroyable"`
}

type WorkspaceCreateResponseResourceAttributesPermissions struct {
	CanCreateStateVersions  bool `json:"can-create-state-versions"`
	CanDestroy              bool `json:"can-destroy"`
	CanForceUnlock          bool `json:"can-force-unlock"`
	CanLock                 bool `json:"can-lock"`
	CanManageRunTasks       bool `json:"can-manage-run-tasks"`
	CanManageTags           bool `json:"can-manage-tags"`
	CanQueueApply           bool `json:"can-queue-apply"`
	CanQueueDestroy         bool `json:"can-queue-destroy"`
	CanQueueRun             bool `json:"can-queue-run"`
	CanReadSettings         bool `json:"can-read-settings"`
	CanReadStateVersions    bool `json:"can-read-state-versions"`
	CanReadVariable         bool `json:"can-read-variable"`
	CanUnlock               bool `json:"can-unlock"`
	CanUpdate               bool `json:"can-update"`
	CanUpdateVariable       bool `json:"can-update-variable"`
	CanReadAssessmentResult bool `json:"can-read-assessment-result"`
	CanForceDelete          bool `json:"can-force-delete"`
}

type WorkspaceCreateResponseResourceAttributesSettingOverwrites struct {
	ExecutionMode bool `json:"execution-mode"`
	AgentPool     bool `json:"agent-pool"`
}

func (o *WorkspaceCreateResponse) MarshalJSON() ([]byte, error) {
	type Alias WorkspaceCreateResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DefaultWorkspaces struct{}

func (o *DefaultWorkspaces) Read(ctx context.Context, organization string, workspace string) (*WorkspaceReadResponse, error) {
	return &WorkspaceReadResponse{
		Data: &jsonapi.Resource[*WorkspaceReadResponseResourceAttributes]{
			ID:         workspace,
			Type:       "workspaces",
			Attributes: &WorkspaceReadResponseResourceAttributes{},
		},
	}, nil
}

func (o *DefaultWorkspaces) Create(ctx context.Context, req *WorkspaceCreateRequest) (*WorkspaceCreateResponse, error) {
	return nil, ErrNotImplemented
}
