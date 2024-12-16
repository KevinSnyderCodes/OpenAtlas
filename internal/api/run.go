package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

// TODO: Implement Runs API
// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/run

var _ Runs = (*DefaultRuns)(nil)

type Runs interface {
	Read(ctx context.Context, runID string) (*RunDocument, error)
	List(ctx context.Context, workspaceID string) (*RunListResponse, error)
	Create(ctx context.Context, req *RunCreateRequest) (*RunDocument, error)
}

type RunDocument jsonapi.Document[
	*jsonapi.Resource[*RunResourceAttributes],
	RunResourceAttributes,
]

func (o *RunDocument) MarshalJSON() ([]byte, error) {
	type Alias RunDocument

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type RunListResponse jsonapi.Document[
	[]*jsonapi.Resource[*RunResourceAttributes],
	RunResourceAttributes,
]

func (o *RunListResponse) MarshalJSON() ([]byte, error) {
	type Alias RunListResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type RunResourceAttributes struct {
	Actions               *RunResourceAttributesActions          `json:"actions"`
	CanceledAt            any                                    `json:"canceled-at"`
	CreatedAt             time.Time                              `json:"created-at"`
	HasChanges            bool                                   `json:"has-changes"`
	AutoApply             bool                                   `json:"auto-apply"`
	AllowEmptyApply       bool                                   `json:"allow-empty-apply"`
	AllowConfigGeneration bool                                   `json:"allow-config-generation"`
	IsDestroy             bool                                   `json:"is-destroy"`
	Message               string                                 `json:"message"`
	PlanOnly              bool                                   `json:"plan-only"`
	Source                string                                 `json:"source"`
	StatusTimestamps      *RunResourceAttributesStatusTimestamps `json:"status-timestamps"`
	Status                string                                 `json:"status"`
	TriggerReason         string                                 `json:"trigger-reason"`
	TargetAddrs           any                                    `json:"target-addrs"`
	Permissions           *RunResourceAttributesPermissions      `json:"permissions"`
	Refresh               bool                                   `json:"refresh"`
	RefreshOnly           bool                                   `json:"refresh-only"`
	ReplaceAddrs          any                                    `json:"replace-addrs"`
	SavePlan              bool                                   `json:"save-plan"`
	Variables             []any                                  `json:"variables"`
}

type RunResourceAttributesActions struct {
	IsCancelable      bool `json:"is-cancelable"`
	IsConfirmable     bool `json:"is-confirmable"`
	IsDiscardable     bool `json:"is-discardable"`
	IsForceCancelable bool `json:"is-force-cancelable"`
}

type RunResourceAttributesStatusTimestamps struct {
	PlanQueueableAt TimeRFC3339Plus `json:"plan-queueable-at"`
}

type RunResourceAttributesPermissions struct {
	CanApply               bool `json:"can-apply"`
	CanCancel              bool `json:"can-cancel"`
	CanComment             bool `json:"can-comment"`
	CanDiscard             bool `json:"can-discard"`
	CanForceExecute        bool `json:"can-force-execute"`
	CanForceCancel         bool `json:"can-force-cancel"`
	CanOverridePolicyCheck bool `json:"can-override-policy-check"`
}

type RunCreateRequest jsonapi.Document[
	*jsonapi.Resource[*RunCreateRequestResourceAttributes],
	RunCreateRequestResourceAttributes,
]

func (o *RunCreateRequest) ConfigurationVersionID() string {
	data, ok := (o.Data.Relationships["configuration-version"].Data).(map[string]any)
	if !ok {
		return ""
	}

	v, ok := data["id"].(string)
	if !ok {
		return ""
	}

	return v
}

type RunCreateRequestResourceAttributes struct {
	AllowEmptyApply       bool     `json:"allow-empty-apply"`
	AllowConfigGeneration bool     `json:"allow-config-generation"`
	AutoApply             bool     `json:"auto-apply"`
	IsDestroy             bool     `json:"is-destroy"`
	Message               string   `json:"message"`
	Refresh               bool     `json:"refresh"`
	RefreshOnly           bool     `json:"refresh-only"`
	ReplaceAddrs          []string `json:"replace-addrs"`
	TargetAddrs           []string `json:"target-addrs"`
	Variables             any      `json:"variables"`
	PlanOnly              bool     `json:"plan-only"`
	SavePlan              bool     `json:"save-plan"`
	TerraformVersion      string   `json:"terraform-version"`
}

func (o *RunCreateRequest) UnmarshalJSON(data []byte) error {
	type Alias RunCreateRequest

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	*o = RunCreateRequest(alias)

	return nil
}

type TFERunDB interface {
	CreateTFERun(ctx context.Context, arg db.CreateTFERunParams) (db.TFERun, error)
	GetTFEPlanByRunID(ctx context.Context, runID string) (db.TFEPlan, error)
	GetTFERun(ctx context.Context, id string) (db.TFERun, error)
	ListTFERuns(ctx context.Context) ([]db.TFERun, error)
}

type TFERun db.TFERun

func (o TFERun) RunExternalID() StrongExternalID {
	return RunInternalID(o.ID).ExternalID()
}

func (o TFERun) ConfigurationVersionExternalID() StrongExternalID {
	return ConfigurationVersionInternalID(o.ConfigurationVersionID).ExternalID()
}

func (o TFERun) RunResource() *jsonapi.Resource[*RunResourceAttributes] {
	return &jsonapi.Resource[*RunResourceAttributes]{
		ID:   o.RunExternalID().String(),
		Type: "runs",
		Attributes: &RunResourceAttributes{
			Status: string(o.Status),
		},
		Relationships: jsonapi.Relationships{
			"configuration-version": &jsonapi.Relationship{
				Data: map[string]string{
					"type": "configuration-versions",
					"id":   o.ConfigurationVersionExternalID().String(),
				},
			},
			"plan": &jsonapi.Relationship{
				Data: map[string]string{
					"type": "plans",
					"id":   "", // TODO: Populate
				},
			},
		},
	}
}

func (o TFERun) RunDocument() *RunDocument {
	return &RunDocument{
		Data: o.RunResource(),
	}
}

type TFERuns []db.TFERun

func (o TFERuns) RunListResponse() *RunListResponse {
	resources := make([]*jsonapi.Resource[*RunResourceAttributes], len(o))

	for i, row := range o {
		resources[i] = (TFERun)(row).RunResource()
	}

	return &RunListResponse{
		Data: resources,
	}
}

type DefaultRuns struct {
	db TFERunDB
}

func NewDefaultRuns(db TFERunDB) *DefaultRuns {
	return &DefaultRuns{
		db: db,
	}
}

func (o *DefaultRuns) Read(ctx context.Context, runID string) (*RunDocument, error) {
	runExternalID := RunExternalID(runID)
	if err := runExternalID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating run id: %w", err)
	}

	runInternalID := runExternalID.InternalID().String()

	row, err := o.db.GetTFERun(ctx, runInternalID)
	if err != nil {
		return nil, fmt.Errorf("error getting run by external id: %w", err)
	}

	resp := (TFERun)(row).RunDocument()

	{
		row, err := o.db.GetTFEPlanByRunID(ctx, runInternalID)
		if err != nil {
			return nil, fmt.Errorf("error getting plan by run id: %w", err)
		}

		resp.Data.Relationships["plan"].Data = map[string]string{
			"type": "plans",
			"id":   PlanInternalID(row.ID).ExternalID().String(),
		}
	}

	return resp, nil
}

func (o *DefaultRuns) List(ctx context.Context, workspaceID string) (*RunListResponse, error) {
	workspaceExternalID := WorkspaceExternalID(workspaceID)
	if err := workspaceExternalID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating workspace id: %w", err)
	}

	if workspaceID != defaultWorkspaceID {
		return nil, fmt.Errorf("workspace not found: %s", workspaceID)
	}

	rows, err := o.db.ListTFERuns(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing runs: %w", err)
	}

	return (TFERuns)(rows).RunListResponse(), nil
}

func (o *DefaultRuns) Create(ctx context.Context, req *RunCreateRequest) (*RunDocument, error) {
	configurationVersionID := req.ConfigurationVersionID()

	configurationVersionExternalID := ConfigurationVersionExternalID(configurationVersionID)
	if err := configurationVersionExternalID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating configuration version id: %w", err)
	}

	arg := db.CreateTFERunParams{
		ID:                     GenerateID(),
		ConfigurationVersionID: configurationVersionExternalID.InternalID().String(),
		Status:                 db.TfeRunStatusPending,
	}
	row, err := o.db.CreateTFERun(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("error creating run: %w", err)
	}

	resp := (TFERun)(row).RunDocument()

	return resp, nil
}
