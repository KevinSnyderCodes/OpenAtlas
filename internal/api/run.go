package api

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/id"
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

func (o *RunCreateRequest) GetConfigurationVersionExternalID() string {
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

func (o TFERun) GetRunID() (id.RunID, error) {
	v, err := id.NewRunIDFromInternalID(o.ID)
	if err != nil {
		return v, fmt.Errorf("error creating run id: %w", err)
	}

	return v, nil
}

func (o TFERun) GetConfigurationVersionID() (id.ConfigurationVersionID, error) {
	v, err := id.NewConfigurationVersionIDFromInternalID(o.ConfigurationVersionID)
	if err != nil {
		return v, fmt.Errorf("error creating configuration version id: %w", err)
	}

	return v, nil
}

func (o TFERun) GetRunResource() (*jsonapi.Resource[*RunResourceAttributes], error) {
	runID, err := o.GetRunID()
	if err != nil {
		return nil, fmt.Errorf("error getting run id: %w", err)
	}

	configurationVersionID, err := o.GetConfigurationVersionID()
	if err != nil {
		return nil, fmt.Errorf("error getting configuration version id: %w", err)
	}

	return &jsonapi.Resource[*RunResourceAttributes]{
		ID:   runID.ExternalID(),
		Type: "runs",
		Attributes: &RunResourceAttributes{
			Status: string(o.Status),
		},
		Relationships: jsonapi.Relationships{
			"configuration-version": &jsonapi.Relationship{
				Data: map[string]string{
					"type": "configuration-versions",
					"id":   configurationVersionID.ExternalID(),
				},
			},
			"plan": &jsonapi.Relationship{
				Data: map[string]string{
					"type": "plans",
					"id":   "", // TODO: Populate
				},
			},
		},
	}, nil
}

func (o TFERun) GetRunDocument() (*RunDocument, error) {
	data, err := o.GetRunResource()
	if err != nil {
		return nil, fmt.Errorf("error getting run resource: %w", err)
	}

	return &RunDocument{
		Data: data,
	}, nil
}

type TFERuns []db.TFERun

func (o TFERuns) GetRunListResponse() (*RunListResponse, error) {
	resources := make([]*jsonapi.Resource[*RunResourceAttributes], len(o))

	for i, row := range o {
		item, err := (TFERun)(row).GetRunResource()
		if err != nil {
			return nil, fmt.Errorf("error getting run resource: %w", err)
		}

		resources[i] = item
	}

	return &RunListResponse{
		Data: resources,
	}, nil
}

type DefaultRuns struct {
	db TFERunDB
}

func NewDefaultRuns(db TFERunDB) *DefaultRuns {
	return &DefaultRuns{
		db: db,
	}
}

func (o *DefaultRuns) Read(ctx context.Context, runExternalID string) (*RunDocument, error) {
	runID, err := id.NewRunIDFromExternalID(runExternalID)
	if err != nil {
		return nil, fmt.Errorf("error creating run id: %w", err)
	}

	runInternalID := runID.InternalID()

	row, err := o.db.GetTFERun(ctx, runInternalID)
	if err != nil {
		return nil, fmt.Errorf("error getting run: %w", err)
	}

	resp, err := (TFERun)(row).GetRunDocument()
	if err != nil {
		return nil, fmt.Errorf("error getting run document: %w", err)
	}

	{
		row, err := o.db.GetTFEPlanByRunID(ctx, runInternalID)
		if err != nil {
			return nil, fmt.Errorf("error getting plan by run id: %w", err)
		}

		planID, err := id.NewPlanIDFromInternalID(row.ID)
		if err != nil {
			return nil, fmt.Errorf("error creating plan id: %w", err)
		}

		resp.Data.Relationships["plan"].Data = map[string]string{
			"type": "plans",
			"id":   planID.ExternalID(),
		}
	}

	return resp, nil
}

func (o *DefaultRuns) List(ctx context.Context, workspaceExternalID string) (*RunListResponse, error) {
	workspaceID, err := id.NewWorkspaceIDFromExternalID(workspaceExternalID)
	if err != nil {
		return nil, fmt.Errorf("error creating workspace id: %w", err)
	}

	if workspaceID.ExternalID() != defaultWorkspaceID {
		return nil, fmt.Errorf("workspace not found: %s", workspaceID)
	}

	rows, err := o.db.ListTFERuns(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing runs: %w", err)
	}

	resp, err := (TFERuns)(rows).GetRunListResponse()
	if err != nil {
		return nil, fmt.Errorf("error getting run list response: %w", err)
	}

	return resp, nil
}

func (o *DefaultRuns) Create(ctx context.Context, req *RunCreateRequest) (*RunDocument, error) {
	configurationVersionExternalID := req.GetConfigurationVersionExternalID()

	configurationVersionID, err := id.NewConfigurationVersionIDFromExternalID(configurationVersionExternalID)
	if err != nil {
		return nil, fmt.Errorf("error creating configuration version id: %w", err)
	}

	arg := db.CreateTFERunParams{
		ID:                     GenerateID(),
		ConfigurationVersionID: configurationVersionID.InternalID(),
		Status:                 db.TfeRunStatusPending,
	}
	row, err := o.db.CreateTFERun(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("error creating run: %w", err)
	}

	resp, err := (TFERun)(row).GetRunDocument()
	if err != nil {
		return nil, fmt.Errorf("error getting run document: %w", err)
	}

	return resp, nil
}
