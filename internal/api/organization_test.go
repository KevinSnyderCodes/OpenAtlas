package api

import (
	"os"
	"testing"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/jsonapi"
	"github.com/stretchr/testify/require"
)

var testOrganizationReadEntitlementsResponseMarshalJSONStruct = OrganizationReadEntitlementsResponse{
	Data: &jsonapi.Resource[*OrganizationReadEntitlementsResponseResourceAttributes]{
		ID:   "org-Bzyc2JuegvVLAibn",
		Type: "entitlement-sets",
		Attributes: &OrganizationReadEntitlementsResponseResourceAttributes{
			CostEstimation:                   false,
			ConfigurationDesigner:            true,
			ModuleTestsGeneration:            false,
			Operations:                       true,
			PrivateModuleRegistry:            true,
			PolicyEnforcement:                false,
			Sentinel:                         false,
			RunTasks:                         false,
			StateStorage:                     true,
			Teams:                            false,
			VCSIntegrations:                  true,
			UsageReporting:                   false,
			UserLimit:                        5,
			SelfServeBilling:                 true,
			AuditLogging:                     false,
			Agents:                           false,
			SSO:                              false,
			RunTaskLimit:                     1,
			RunTaskWorkspaceLimit:            10,
			RunTaskMandatoryEnforcementLimit: 1,
			PolicySetLimit:                   1,
			PolicyLimit:                      5,
			PolicyMandatoryEnforcementLimit:  nil,
			VersionedPolicySetLimit:          nil,
		},
		Links: jsonapi.Links{
			"self": "/api/v2/entitlement-sets/org-Bzyc2JuegvVLAibn",
		},
	},
}

func TestOrganizationReadEntitlementsResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		o        *OrganizationReadEntitlementsResponse
		wantFile string
		wantErr  bool
	}{
		{
			name:     "success",
			o:        &testOrganizationReadEntitlementsResponseMarshalJSONStruct,
			wantFile: "./testdata/organization_readentitlements_response.json",
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
				t.Errorf("OrganizationReadEntitlementsResponse.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(want), string(got))
		})
	}
}
