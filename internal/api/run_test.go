package api

import (
	"os"
	"testing"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
	"github.com/stretchr/testify/require"
)

var testRunReadResponseMarshalJSONStruct = RunReadResponse{
	Data: &jsonapi.Resource[*RunReadResponseResourceAttributes]{
		ID:   "run-CZcmD7eagjhyX0vN",
		Type: "runs",
		Attributes: &RunReadResponseResourceAttributes{
			Actions: &RunReadResponseResourceAttributesActions{
				IsCancelable:      true,
				IsConfirmable:     false,
				IsDiscardable:     false,
				IsForceCancelable: false,
			},
			CanceledAt:            nil,
			CreatedAt:             time.Date(2021, 5, 24, 7, 38, 4, 171000000, time.UTC),
			HasChanges:            false,
			AutoApply:             false,
			AllowEmptyApply:       false,
			AllowConfigGeneration: false,
			IsDestroy:             false,
			Message:               "Custom message",
			PlanOnly:              false,
			Source:                "tfe-api",
			StatusTimestamps: &RunReadResponseResourceAttributesStatusTimestamps{
				PlanQueueableAt: TimeRFC3339Plus(time.Date(2021, 5, 24, 7, 38, 4, 0, time.UTC)),
			},
			Status:        "pending",
			TriggerReason: "manual",
			TargetAddrs:   nil,
			Permissions: &RunReadResponseResourceAttributesPermissions{
				CanApply:               true,
				CanCancel:              true,
				CanComment:             true,
				CanDiscard:             true,
				CanForceExecute:        true,
				CanForceCancel:         true,
				CanOverridePolicyCheck: true,
			},
			Refresh:      false,
			RefreshOnly:  false,
			ReplaceAddrs: nil,
			SavePlan:     false,
			Variables:    []any{},
		},
		Relationships: jsonapi.Relationships{
			"apply":                 &jsonapi.Relationship{},
			"comments":              &jsonapi.Relationship{},
			"configuration-version": &jsonapi.Relationship{},
			"cost-estimate":         &jsonapi.Relationship{},
			"created-by":            &jsonapi.Relationship{},
			"input-state-version":   &jsonapi.Relationship{},
			"plan":                  &jsonapi.Relationship{},
			"run-events":            &jsonapi.Relationship{},
			"policy-checks":         &jsonapi.Relationship{},
			"task-stages":           &jsonapi.Relationship{},
			"workspace":             &jsonapi.Relationship{},
			"workspace-run-alerts":  &jsonapi.Relationship{},
		},
		Links: jsonapi.Links{
			"self": "/api/v2/runs/run-bWSq4YeYpfrW4mx7",
		},
	},
}

var testRunCreateRequestUnmarshalJSONStruct = RunCreateRequest{
	Data: &jsonapi.Resource[*RunCreateRequestResourceAttributes]{
		Attributes: &RunCreateRequestResourceAttributes{
			Message: "Custom message",
		},
		Type: "runs",
		Relationships: jsonapi.Relationships{
			"workspace": &jsonapi.Relationship{
				Data: map[string]any{
					"type": "workspaces",
					"id":   "ws-LLGHCr4SWy28wyGN",
				},
			},
			"configuration-version": &jsonapi.Relationship{
				Data: map[string]any{
					"type": "configuration-versions",
					"id":   "cv-n4XQPBa2QnecZJ4G",
				},
			},
		},
	},
}

var testRunCreateResponseMarshalJSONStruct = RunCreateResponse{
	Data: &jsonapi.Resource[*RunCreateResponseResourceAttributes]{
		ID:   "run-CZcmD7eagjhyX0vN",
		Type: "runs",
		Attributes: &RunCreateResponseResourceAttributes{
			Actions: &RunCreateResponseResourceAttributesActions{
				IsCancelable:      true,
				IsConfirmable:     false,
				IsDiscardable:     false,
				IsForceCancelable: false,
			},
			CanceledAt:            nil,
			CreatedAt:             time.Date(2021, 5, 24, 7, 38, 4, 171000000, time.UTC),
			HasChanges:            false,
			AutoApply:             false,
			AllowEmptyApply:       false,
			AllowConfigGeneration: false,
			IsDestroy:             false,
			Message:               "Custom message",
			PlanOnly:              false,
			Source:                "tfe-api",
			StatusTimestamps: &RunCreateResponseResourceAttributesStatusTimestamps{
				PlanQueueableAt: TimeRFC3339Plus(time.Date(2021, 5, 24, 7, 38, 4, 0, time.UTC)),
			},
			Status:        "pending",
			TriggerReason: "manual",
			TargetAddrs:   nil,
			Permissions: &RunCreateResponseResourceAttributesPermissions{
				CanApply:               true,
				CanCancel:              true,
				CanComment:             true,
				CanDiscard:             true,
				CanForceExecute:        true,
				CanForceCancel:         true,
				CanOverridePolicyCheck: true,
			},
			Refresh:      false,
			RefreshOnly:  false,
			ReplaceAddrs: nil,
			SavePlan:     false,
			Variables:    []any{},
		},
		Relationships: jsonapi.Relationships{
			"apply":                 &jsonapi.Relationship{},
			"comments":              &jsonapi.Relationship{},
			"configuration-version": &jsonapi.Relationship{},
			"cost-estimate":         &jsonapi.Relationship{},
			"created-by":            &jsonapi.Relationship{},
			"input-state-version":   &jsonapi.Relationship{},
			"plan":                  &jsonapi.Relationship{},
			"run-events":            &jsonapi.Relationship{},
			"policy-checks":         &jsonapi.Relationship{},
			"workspace":             &jsonapi.Relationship{},
			"workspace-run-alerts":  &jsonapi.Relationship{},
		},
		Links: jsonapi.Links{
			"self": "/api/v2/runs/run-CZcmD7eagjhyX0vN",
		},
	},
}

func TestRunReadResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *RunReadResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testRunReadResponseMarshalJSONStruct,
			wantFile: "./testdata/run_read_response.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := os.ReadFile(tt.wantFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			got, err := tt.o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("RunReadResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}

func TestRunCreateRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		dataFile string
		want     *RunCreateRequest
		wantErr  bool
	}{
		{
			name:     "success",
			dataFile: "./testdata/run_create_request.json",
			want:     &testRunCreateRequestUnmarshalJSONStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.dataFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			o := RunCreateRequest{}
			if err := o.UnmarshalJSON(data); (err != nil) != tt.wantErr {
				t.Errorf("RunCreateRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.want, &o)
		})
	}
}

func TestRunCreateResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *RunCreateResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testRunCreateResponseMarshalJSONStruct,
			wantFile: "./testdata/run_create_response.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := os.ReadFile(tt.wantFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			got, err := tt.o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("RunCreateResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}
