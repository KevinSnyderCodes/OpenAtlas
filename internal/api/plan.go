package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

// TODO: Implement Plans API
// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/plans

var _ Plans = (*DefaultPlans)(nil)

type Plans interface {
	Read(ctx context.Context, planID string) (*PlanDocument, error)
}

type PlanDocument jsonapi.Document[
	*jsonapi.Resource[*PlanResourceAttributes],
	PlanResourceAttributes,
]

func (o *PlanDocument) MarshalJSON() ([]byte, error) {
	type Alias PlanDocument

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type PlanResourceAttributes struct {
	ExecutionDetails       *PlanResourceAttributesExecutionDetails `json:"execution-details"`
	GeneratedConfiguration bool                                    `json:"generated-configuration"`
	HasChanges             bool                                    `json:"has-changes"`
	ResourceAdditions      int                                     `json:"resource-additions"`
	ResourceChanges        int                                     `json:"resource-changes"`
	ResourceDestructions   int                                     `json:"resource-destructions"`
	ResourceImports        int                                     `json:"resource-imports"`
	Status                 string                                  `json:"status"`
	StatusTimestamps       *PlanResourceAttributesStatusTimestamps `json:"status-timestamps"`
	LogReadURL             string                                  `json:"log-read-url"`
}

type PlanResourceAttributesExecutionDetails struct {
	Mode string `json:"mode"`
}

type PlanResourceAttributesStatusTimestamps struct {
	QueuedAt   TimeRFC3339Plus `json:"queued-at"`
	PendingAt  TimeRFC3339Plus `json:"pending-at"`
	StartedAt  TimeRFC3339Plus `json:"started-at"`
	FinishedAt TimeRFC3339Plus `json:"finished-at"`
}

type TFEPlanDB interface {
	CreateTFEPlan(ctx context.Context, arg db.CreateTFEPlanParams) (db.TFEPlan, error)
	GetTFEPlan(ctx context.Context, id string) (db.TFEPlan, error)
}

type TFEPlan db.TFEPlan

func (o TFEPlan) PlanExternalID() PlanExternalID {
	return PlanExternalID(o.ID)
}

func (o TFEPlan) RunExternalID() RunExternalID {
	return RunExternalID(o.RunID)
}

func (o TFEPlan) PlanResource() *jsonapi.Resource[*PlanResourceAttributes] {
	return &jsonapi.Resource[*PlanResourceAttributes]{
		ID:   o.PlanExternalID().String(),
		Type: "plans",
		Attributes: &PlanResourceAttributes{
			Status:     string(o.Status),
			LogReadURL: o.LogReadUrl,
		},
		Relationships: jsonapi.Relationships{
			"run": &jsonapi.Relationship{
				Data: map[string]string{
					"type": "runs",
					"id":   o.RunExternalID().String(),
				},
			},
		},
	}
}

func (o TFEPlan) PlanDocument() *PlanDocument {
	return &PlanDocument{
		Data: o.PlanResource(),
	}
}

type DefaultPlans struct {
	db TFEPlanDB
}

func NewDefaultPlans(db TFEPlanDB) *DefaultPlans {
	return &DefaultPlans{
		db: db,
	}
}

func (o *DefaultPlans) Read(ctx context.Context, planID string) (*PlanDocument, error) {
	planExternalID := PlanExternalID(planID)
	if err := planExternalID.Validate(); err != nil {
		return nil, fmt.Errorf("error validating plan id: %w", err)
	}

	planInternalID := planExternalID.InternalID().String()

	row, err := o.db.GetTFEPlan(ctx, planInternalID)
	if err != nil {
		return nil, fmt.Errorf("error getting tfe plan: %w", err)
	}

	resp := (TFEPlan)(row).PlanDocument()

	return resp, nil
}
