package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

// TODO: Implement Runs API
// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/run

var _ Runs = (*DefaultRuns)(nil)

type Runs interface {
	Read(ctx context.Context, req *RunReadRequest) (*RunReadResponse, error)
	Create(ctx context.Context, req *RunCreateRequest) (*RunCreateResponse, error)
}

type RunReadRequest struct {
	// TODO: Populate
}

type RunReadResponse jsonapi.Document[
	*jsonapi.Resource[*RunReadResponseResourceAttributes],
	RunReadResponseResourceAttributes,
]

type RunReadResponseResourceAttributes struct {
	Actions               *RunReadResponseResourceAttributesActions          `json:"actions"`
	CanceledAt            any                                                `json:"canceled-at"`
	CreatedAt             time.Time                                          `json:"created-at"`
	HasChanges            bool                                               `json:"has-changes"`
	AutoApply             bool                                               `json:"auto-apply"`
	AllowEmptyApply       bool                                               `json:"allow-empty-apply"`
	AllowConfigGeneration bool                                               `json:"allow-config-generation"`
	IsDestroy             bool                                               `json:"is-destroy"`
	Message               string                                             `json:"message"`
	PlanOnly              bool                                               `json:"plan-only"`
	Source                string                                             `json:"source"`
	StatusTimestamps      *RunReadResponseResourceAttributesStatusTimestamps `json:"status-timestamps"`
	Status                string                                             `json:"status"`
	TriggerReason         string                                             `json:"trigger-reason"`
	TargetAddrs           any                                                `json:"target-addrs"`
	Permissions           *RunReadResponseResourceAttributesPermissions      `json:"permissions"`
	Refresh               bool                                               `json:"refresh"`
	RefreshOnly           bool                                               `json:"refresh-only"`
	ReplaceAddrs          any                                                `json:"replace-addrs"`
	SavePlan              bool                                               `json:"save-plan"`
	Variables             []any                                              `json:"variables"`
}

type RunReadResponseResourceAttributesActions struct {
	IsCancelable      bool `json:"is-cancelable"`
	IsConfirmable     bool `json:"is-confirmable"`
	IsDiscardable     bool `json:"is-discardable"`
	IsForceCancelable bool `json:"is-force-cancelable"`
}

type RunReadResponseResourceAttributesStatusTimestamps struct {
	PlanQueueableAt TimeRFC3339Plus `json:"plan-queueable-at"`
}

type RunReadResponseResourceAttributesPermissions struct {
	CanApply               bool `json:"can-apply"`
	CanCancel              bool `json:"can-cancel"`
	CanComment             bool `json:"can-comment"`
	CanDiscard             bool `json:"can-discard"`
	CanForceExecute        bool `json:"can-force-execute"`
	CanForceCancel         bool `json:"can-force-cancel"`
	CanOverridePolicyCheck bool `json:"can-override-policy-check"`
}

func (o *RunReadResponse) MarshalJSON() ([]byte, error) {
	type Alias RunReadResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type RunCreateRequest jsonapi.Document[
	*jsonapi.Resource[*RunCreateRequestResourceAttributes],
	RunCreateRequestResourceAttributes,
]

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

type RunCreateResponse jsonapi.Document[
	*jsonapi.Resource[*RunCreateResponseResourceAttributes],
	RunCreateResponseResourceAttributes,
]

type RunCreateResponseResourceAttributes struct {
	Actions               *RunCreateResponseResourceAttributesActions          `json:"actions"`
	CanceledAt            any                                                  `json:"canceled-at"`
	CreatedAt             time.Time                                            `json:"created-at"`
	HasChanges            bool                                                 `json:"has-changes"`
	AutoApply             bool                                                 `json:"auto-apply"`
	AllowEmptyApply       bool                                                 `json:"allow-empty-apply"`
	AllowConfigGeneration bool                                                 `json:"allow-config-generation"`
	IsDestroy             bool                                                 `json:"is-destroy"`
	Message               string                                               `json:"message"`
	PlanOnly              bool                                                 `json:"plan-only"`
	Source                string                                               `json:"source"`
	StatusTimestamps      *RunCreateResponseResourceAttributesStatusTimestamps `json:"status-timestamps"`
	Status                string                                               `json:"status"`
	TriggerReason         string                                               `json:"trigger-reason"`
	TargetAddrs           any                                                  `json:"target-addrs"`
	Permissions           *RunCreateResponseResourceAttributesPermissions      `json:"permissions"`
	Refresh               bool                                                 `json:"refresh"`
	RefreshOnly           bool                                                 `json:"refresh-only"`
	ReplaceAddrs          any                                                  `json:"replace-addrs"`
	SavePlan              bool                                                 `json:"save-plan"`
	Variables             []any                                                `json:"variables"`
}

type RunCreateResponseResourceAttributesActions struct {
	IsCancelable      bool `json:"is-cancelable"`
	IsConfirmable     bool `json:"is-confirmable"`
	IsDiscardable     bool `json:"is-discardable"`
	IsForceCancelable bool `json:"is-force-cancelable"`
}

type RunCreateResponseResourceAttributesStatusTimestamps struct {
	PlanQueueableAt TimeRFC3339Plus `json:"plan-queueable-at"`
}

type RunCreateResponseResourceAttributesPermissions struct {
	CanApply               bool `json:"can-apply"`
	CanCancel              bool `json:"can-cancel"`
	CanComment             bool `json:"can-comment"`
	CanDiscard             bool `json:"can-discard"`
	CanForceExecute        bool `json:"can-force-execute"`
	CanForceCancel         bool `json:"can-force-cancel"`
	CanOverridePolicyCheck bool `json:"can-override-policy-check"`
}

func (o *RunCreateResponse) MarshalJSON() ([]byte, error) {
	type Alias RunCreateResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DefaultRuns struct{}

func (o *DefaultRuns) Read(ctx context.Context, req *RunReadRequest) (*RunReadResponse, error) {
	return nil, ErrNotImplemented
}

func (o *DefaultRuns) Create(ctx context.Context, req *RunCreateRequest) (*RunCreateResponse, error) {
	return nil, ErrNotImplemented
}
