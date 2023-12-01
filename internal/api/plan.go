package api

import (
	"context"
	"encoding/json"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
)

// TODO: Implement Plans API
// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/plans

var _ Plans = (*DefaultPlans)(nil)

type Plans interface {
	Read(ctx context.Context, req *PlanReadRequest) (*PlanReadResponse, error)
}

type PlanReadRequest struct {
	// TODO: Populate
}

type PlanReadResponse jsonapi.Document[
	*jsonapi.Resource[*PlanReadResponseResourceAttributes],
	PlanReadResponseResourceAttributes,
]

type PlanReadResponseResourceAttributes struct {
	ExecutionDetails       *PlanReadResponseResourceAttributesExecutionDetails `json:"execution-details"`
	GeneratedConfiguration bool                                                `json:"generated-configuration"`
	HasChanges             bool                                                `json:"has-changes"`
	ResourceAdditions      int                                                 `json:"resource-additions"`
	ResourceChanges        int                                                 `json:"resource-changes"`
	ResourceDestructions   int                                                 `json:"resource-destructions"`
	ResourceImports        int                                                 `json:"resource-imports"`
	Status                 string                                              `json:"status"`
	StatusTimestamps       *PlanReadResponseResourceAttributesStatusTimestamps `json:"status-timestamps"`
	LogReadURL             string                                              `json:"log-read-url"`
}

type PlanReadResponseResourceAttributesExecutionDetails struct {
	Mode string `json:"mode"`
}

type PlanReadResponseResourceAttributesStatusTimestamps struct {
	QueuedAt   TimeRFC3339Plus `json:"queued-at"`
	PendingAt  TimeRFC3339Plus `json:"pending-at"`
	StartedAt  TimeRFC3339Plus `json:"started-at"`
	FinishedAt TimeRFC3339Plus `json:"finished-at"`
}

func (o *PlanReadResponse) MarshalJSON() ([]byte, error) {
	type Alias PlanReadResponse

	data, err := json.Marshal((*Alias)(o))
	if err != nil {
		return nil, err
	}

	return data, nil
}

type DefaultPlans struct{}

func (o *DefaultPlans) Read(ctx context.Context, req *PlanReadRequest) (*PlanReadResponse, error) {
	return nil, ErrNotImplemented
}
