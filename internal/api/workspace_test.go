package api

import (
	"os"
	"testing"
	"time"

	xjson "github.com/KevinSnyderCodes/OpenAtlas/internal/x/json"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
	"github.com/stretchr/testify/require"
)

var testWorkspaceReadResponseMarshalJSONStruct = WorkspaceReadResponse{
	Data: &jsonapi.Resource[*WorkspaceReadResponseResourceAttributes]{
		ID:   "ws-qPhan8kDLymzv2uS",
		Type: "workspaces",
		Attributes: &WorkspaceReadResponseResourceAttributes{
			Actions: &WorkspaceReadResponseResourceAttributesActions{
				IsDestroyable: true,
			},
			AllowDestroyPlan:            true,
			ApplyDurationAverage:        158000,
			AutoApply:                   false,
			AutoApplyRunTrigger:         false,
			AutoDestroyAt:               nil,
			AutoDestroyActivityDuration: nil,
			CreatedAt:                   time.Date(2021, 6, 3, 17, 50, 20, 307000000, time.UTC),
			Description:                 "An example workspace for documentation.",
			Environment:                 "default",
			ExecutionMode:               "remote",
			FileTriggersEnabled:         true,
			GlobalRemoteState:           false,
			LatestChangeAt:              time.Date(2021, 6, 23, 17, 50, 48, 815000000, time.UTC),
			Locked:                      false,
			Name:                        "workspace-1",
			Operations:                  true,
			Permissions: &WorkspaceReadResponseResourceAttributesPermissions{
				CanCreateStateVersions: true,
				CanDestroy:             true,
				CanForceUnlock:         true,
				CanLock:                true,
				CanManageRunTasks:      true,
				CanManageTags:          true,
				CanQueueApply:          true,
				CanQueueDestroy:        true,
				CanQueueRun:            true,
				CanReadSettings:        true,
				CanReadStateVersions:   true,
				CanReadVariable:        true,
				CanUnlock:              true,
				CanUpdate:              true,
				CanUpdateVariable:      true,
				CanForceDelete:         true,
			},
			PlanDurationAverage:        20000,
			PolicyCheckFailures:        nil,
			QueueAllRuns:               false,
			ResourceCount:              0,
			RunFailures:                6,
			Source:                     "terraform",
			SourceName:                 nil,
			SourceURL:                  nil,
			SpeculativeEnabled:         true,
			StructuredRunOutputEnabled: false,
			TerraformVersion:           "0.15.3",
			TriggerPrefixes:            []string{},
			UpdatedAt:                  time.Date(2021, 8, 16, 18, 54, 6, 874000000, time.UTC),
			VCSRepo:                    nil,
			VCSRepoIdentifier:          nil,
			WorkingDirectory:           nil,
			WorkspaceKPIsRunsCount:     7,
			SettingOverwrites: &WorkspaceReadResponseResourceAttributesSettingOverwrites{
				ExecutionMode: true,
				AgentPool:     true,
			},
		},
		Links: jsonapi.Links{
			"self": "/api/v2/organizations/my-organization/workspaces/workspace-1",
		},
		Relationships: jsonapi.Relationships{
			"agent-pool": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "apool-QxGd2tRjympfMvQc",
					"type": "agent-pools",
				},
			},
			"current-configuration-version": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "cv-sixaaRuRwutYg5fH",
					"type": "configuration-versions",
				},
				Links: jsonapi.Links{
					"related": "/api/v2/configuration-versions/cv-sixaaRuRwutYg5fH",
				},
			},
			"current-run": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "run-UyCw2TDCmxtfdjmy",
					"type": "runs",
				},
				Links: jsonapi.Links{
					"related": "/api/v2/runs/run-UyCw2TDCmxtfdjmy",
				},
			},
			"current-state-version": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "sv-TAjm2vFZqY396qY6",
					"type": "state-versions",
				},
				Links: jsonapi.Links{
					"related": "/api/v2/workspaces/ws-qPhan8kDLymzv2uS/current-state-version",
				},
			},
			"latest-run": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "run-UyCw2TDCmxtfdjmy",
					"type": "runs",
				},
				Links: jsonapi.Links{
					"related": "/api/v2/runs/run-UyCw2TDCmxtfdjmy",
				},
			},
			"organization": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "my-organization",
					"type": "organizations",
				},
			},
			"project": &jsonapi.Relationship{
				Data: map[string]any{
					"type": "projects",
					"id":   "prj-7HWWPGY3fYxztELU",
				},
			},
			"outputs": &jsonapi.Relationship{
				Data: []any{},
			},
			"readme": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "227247",
					"type": "workspace-readme",
				},
			},
			"remote-state-consumers": &jsonapi.Relationship{
				Links: jsonapi.Links{
					"related": "/api/v2/workspaces/ws-qPhan8kDLymzv2uS/relationships/remote-state-consumers",
				},
			},
		},
	},
}

var testWorkspaceCreateRequestUnmarshalJSONStruct = WorkspaceCreateRequest{
	Data: &jsonapi.Resource[*WorkspaceCreateRequestResourceAttributes]{
		Type: "workspaces",
		Attributes: &WorkspaceCreateRequestResourceAttributes{
			Name:          "workspace-1",
			ResourceCount: 0,
			UpdatedAt:     time.Date(2017, 11, 29, 19, 18, 9, 976000000, time.UTC),
		},
	},
}

var testWorkspaceCreateResponseMarshalJSONStruct = WorkspaceCreateResponse{
	Data: &jsonapi.Resource[*WorkspaceCreateResponseResourceAttributes]{
		ID:   "ws-6jrRyVDv1J8zQMB5",
		Type: "workspaces",
		Attributes: &WorkspaceCreateResponseResourceAttributes{
			Actions: &WorkspaceCreateResponseResourceAttributesActions{
				IsDestroyable: true,
			},
			AllowDestroyPlan:            true,
			ApplyDurationAverage:        nil,
			AutoApply:                   false,
			AutoApplyRunTrigger:         false,
			AutoDestroyAt:               nil,
			AutoDestroyActivityDuration: nil,
			CreatedAt:                   time.Date(2021, 8, 16, 21, 22, 49, 566000000, time.UTC),
			Description:                 nil,
			Environment:                 "default",
			ExecutionMode:               "agent",
			FileTriggersEnabled:         true,
			GlobalRemoteState:           false,
			LatestChangeAt:              time.Date(2021, 8, 16, 21, 22, 49, 566000000, time.UTC),
			LastAssessmentResultAt:      time.Date(2021, 8, 17, 21, 20, 12, 908000000, time.UTC),
			Locked:                      false,
			Name:                        "workspace-1",
			Operations:                  true,
			Permissions: &WorkspaceCreateResponseResourceAttributesPermissions{
				CanCreateStateVersions:  true,
				CanDestroy:              true,
				CanForceUnlock:          true,
				CanLock:                 true,
				CanManageRunTasks:       true,
				CanManageTags:           true,
				CanQueueApply:           true,
				CanQueueDestroy:         true,
				CanQueueRun:             true,
				CanReadSettings:         true,
				CanReadStateVersions:    true,
				CanReadVariable:         true,
				CanUnlock:               true,
				CanUpdate:               true,
				CanUpdateVariable:       true,
				CanReadAssessmentResult: true,
				CanForceDelete:          true,
			},
			PlanDurationAverage:        nil,
			PolicyCheckFailures:        nil,
			QueueAllRuns:               false,
			ResourceCount:              0,
			RunFailures:                nil,
			Source:                     "tfe-api",
			SourceName:                 nil,
			SourceURL:                  nil,
			SpeculativeEnabled:         true,
			StructuredRunOutputEnabled: true,
			TerraformVersion:           "1.0.4",
			TriggerPrefixes:            []string{},
			UpdatedAt:                  time.Date(2021, 8, 16, 21, 22, 49, 566000000, time.UTC),
			VCSRepo:                    nil,
			VCSRepoIdentifier:          nil,
			WorkingDirectory:           nil,
			WorkspaceKPIsRunsCount:     nil,
			AssessmentsEnabled:         false,
			SettingOverwrites: &WorkspaceCreateResponseResourceAttributesSettingOverwrites{
				ExecutionMode: true,
				AgentPool:     true,
			},
		},
		Links: jsonapi.Links{
			"self": "/api/v2/organizations/my-organization/workspaces/workspace-1",
		},
		Relationships: jsonapi.Relationships{
			"agent-pool": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "apool-QxGd2tRjympfMvQc",
					"type": "agent-pools",
				},
			},
			"current-configuration-version": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"current-run": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"current-state-version": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"current-assessment-result": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"latest-run": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"organization": &jsonapi.Relationship{
				Data: map[string]any{
					"id":   "my-organization",
					"type": "organizations",
				},
			},
			"outputs": &jsonapi.Relationship{
				Data: []any{},
			},
			"readme": &jsonapi.Relationship{
				Data: xjson.Null,
			},
			"remote-state-consumers": &jsonapi.Relationship{
				Links: jsonapi.Links{
					"related": "/api/v2/workspaces/ws-6jrRyVDv1J8zQMB5/relationships/remote-state-consumers",
				},
			},
		},
	},
}

func TestWorkspaceReadResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *WorkspaceReadResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testWorkspaceReadResponseMarshalJSONStruct,
			wantFile: "./testdata/workspace_read_response.json",
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
				t.Errorf("WorkspaceReadResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}

func TestWorkspaceCreateRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		dataFile string
		want     *WorkspaceCreateRequest
		wantErr  bool
	}{
		{
			name:     "success without vcs",
			dataFile: "./testdata/workspace_create_request_withoutvcs.json",
			want:     &testWorkspaceCreateRequestUnmarshalJSONStruct,
		},
		// TODO: Add "success with vcs"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.dataFile)
			if err != nil {
				t.Errorf("error reading file: %v", err)
				return
			}

			var o WorkspaceCreateRequest
			if err := o.UnmarshalJSON(data); (err != nil) != tt.wantErr {
				t.Errorf("WorkspaceCreateRequest.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			require.Equal(t, tt.want, &o)
		})
	}
}

func TestWorkspaceCreateResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *WorkspaceCreateResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success without vcs",
			o:        &testWorkspaceCreateResponseMarshalJSONStruct,
			wantFile: "./testdata/workspace_create_response_withoutvcs.json",
		},
		// TODO: Add "success with vcs"
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
				t.Errorf("WorkspaceCreateResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}
