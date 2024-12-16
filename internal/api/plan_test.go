package api

import (
	"os"
	"testing"
	"time"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
	"github.com/stretchr/testify/require"
)

var testPlanReadResponseMarshalJSONStruct = PlanReadResponse{
	Data: &jsonapi.Resource[*PlanResourceAttributes]{
		ID:   "plan-8F5JFydVYAmtTjET",
		Type: "plans",
		Attributes: &PlanResourceAttributes{
			ExecutionDetails: &PlanResourceAttributesExecutionDetails{
				Mode: "remote",
			},
			GeneratedConfiguration: false,
			HasChanges:             true,
			ResourceAdditions:      0,
			ResourceChanges:        1,
			ResourceDestructions:   0,
			ResourceImports:        0,
			Status:                 "finished",
			StatusTimestamps: &PlanResourceAttributesStatusTimestamps{
				QueuedAt:   TimeRFC3339Plus(time.Date(2018, 7, 2, 22, 29, 53, 0, time.UTC)),
				PendingAt:  TimeRFC3339Plus(time.Date(2018, 7, 2, 22, 29, 53, 0, time.UTC)),
				StartedAt:  TimeRFC3339Plus(time.Date(2018, 7, 2, 22, 29, 54, 0, time.UTC)),
				FinishedAt: TimeRFC3339Plus(time.Date(2018, 7, 2, 22, 29, 58, 0, time.UTC)),
			},
			LogReadURL: "https://archivist.terraform.io/v1/object/dmF1bHQ6djE6OFA1eEdlSFVHRSs4YUcwaW83a1dRRDA0U2E3T3FiWk1HM2NyQlNtcS9JS1hHN3dmTXJmaFhEYTlHdTF1ZlgxZ2wzVC9kVTlNcjRPOEJkK050VFI3U3dvS2ZuaUhFSGpVenJVUFYzSFVZQ1VZYno3T3UyYjdDRVRPRE5pbWJDVTIrNllQTENyTndYd1Y0ak1DL1dPVlN1VlNxKzYzbWlIcnJPa2dRRkJZZGtFeTNiaU84YlZ4QWs2QzlLY3VJb3lmWlIrajF4a1hYZTlsWnFYemRkL2pNOG9Zc0ZDakdVMCtURUE3dDNMODRsRnY4cWl1dUN5dUVuUzdnZzFwL3BNeHlwbXNXZWRrUDhXdzhGNnF4c3dqaXlZS29oL3FKakI5dm9uYU5ZKzAybnloREdnQ3J2Rk5WMlBJemZQTg",
		},
		Relationships: jsonapi.Relationships{
			"state-versions": &jsonapi.Relationship{
				Data: []any{},
			},
		},
		Links: jsonapi.Links{
			"self":        "/api/v2/plans/plan-8F5JFydVYAmtTjET",
			"json-output": "/api/v2/plans/plan-8F5JFydVYAmtTjET/json-output",
		},
	},
}

func TestPlanReadResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *PlanReadResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testPlanReadResponseMarshalJSONStruct,
			wantFile: "./testdata/plan_read_response.json",
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
				t.Errorf("PlanReadResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}
