package api

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/require"
)

var testOrganizationListResponseMarshalStruct = OrganizationListResponse{
	Items: []*OrganizationListResponseItem{
		{
			ID:                     "hashicorp",
			ExternalID:             "org-Hysjx5eUviuKVCJY",
			CreatedAt:              time.Date(2021, 8, 24, 23, 10, 4, 675000000, time.UTC),
			Email:                  "hashicorp@example.com",
			SessionTimeout:         nil,
			SessionRemember:        nil,
			CollaboratorAuthPolicy: "password",
			PlanExpired:            false,
			PlanExpiresAt:          nil,
			PlanIsTrial:            false,
			PlanIsEnterprise:       false,
			PlanIdentifier:         "developer",
			CostEstimationEnabled:  true,
			SendPassingStatusesForUntriggeredSpeculativePlans: true,
			AllowForceDeleteWorkspaces:                        true,
			Name:                                              "hashicorp",
			Permissions: &OrganizationListResponseItemPermissions{
				CanUpdate:                true,
				CanDestroy:               true,
				CanAccessViaTeams:        true,
				CanCreateModule:          true,
				CanCreateTeam:            true,
				CanCreateWorkspace:       true,
				CanManageUsers:           true,
				CanManageSubscription:    true,
				CanManageSSO:             true,
				CanUpdateOAuth:           true,
				CanUpdateSentinel:        true,
				CanUpdateSSHKeys:         true,
				CanUpdateAPIToken:        true,
				CanTraverse:              true,
				CanStartTrial:            true,
				CanUpdateAgentPools:      true,
				CanManageTags:            true,
				CanManageVarsets:         true,
				CanReadVarsets:           true,
				CanManagePublicProviders: true,
				CanCreateProvider:        true,
				CanManagePublicModules:   true,
				CanManageCustomProviders: false,
				CanManageRunTasks:        false,
				CanReadRunTasks:          false,
				CanCreateProject:         true,
			},
			FairRunQueuingEnabled:           true,
			SAMLEnabled:                     false,
			OwnersTeamSAMLRoleID:            nil,
			TwoFactorConformant:             false,
			AssessmentsEnforced:             false,
			DefaultExecutionMode:            "remote",
			RelationshipDefaultAgentPool:    nil,
			RelationshipOAuthTokens:         nil,
			RelationshipAuthenticationToken: nil,
			RelationshipEntitlementSet: &OrganizationListResponseItemRelationshipEntitlementSet{
				ID: "org-Hysjx5eUviuKVCJY",
			},
			RelationshipSubscription: nil,
			RelationshipLinksOAuthTokens: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp/oauth-tokens",
			},
			RelationshipLinksAuthenticationToken: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp/authentication-token",
			},
			RelationshipLinksEntitlementSet: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp/entitlement-set",
			},
			RelationshipLinksSubscription: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp/subscription",
			},
			Links: &jsonapi.Links{
				"self": "/api/v2/organizations/hashicorp",
			},
		},
		{
			ID:                     "hashicorp-two",
			ExternalID:             "org-iJ5tr4WgB4WpA1hD",
			CreatedAt:              time.Date(2022, 1, 4, 18, 57, 16, 36000000, time.UTC),
			Email:                  "hashicorp@example.com",
			SessionTimeout:         nil,
			SessionRemember:        nil,
			CollaboratorAuthPolicy: "password",
			PlanExpired:            false,
			PlanExpiresAt:          nil,
			PlanIsTrial:            false,
			PlanIsEnterprise:       false,
			PlanIdentifier:         "free",
			CostEstimationEnabled:  false,
			SendPassingStatusesForUntriggeredSpeculativePlans: false,
			AllowForceDeleteWorkspaces:                        false,
			Name:                                              "hashicorp-two",
			Permissions: &OrganizationListResponseItemPermissions{
				CanUpdate:                true,
				CanDestroy:               true,
				CanAccessViaTeams:        true,
				CanCreateModule:          true,
				CanCreateTeam:            false,
				CanCreateWorkspace:       true,
				CanManageUsers:           true,
				CanManageSubscription:    true,
				CanManageSSO:             false,
				CanUpdateOAuth:           true,
				CanUpdateSentinel:        false,
				CanUpdateSSHKeys:         true,
				CanUpdateAPIToken:        true,
				CanTraverse:              true,
				CanStartTrial:            true,
				CanUpdateAgentPools:      false,
				CanManageTags:            true,
				CanManageVarsets:         true,
				CanReadVarsets:           true,
				CanManagePublicProviders: true,
				CanCreateProvider:        true,
				CanManagePublicModules:   true,
				CanManageCustomProviders: false,
				CanManageRunTasks:        false,
				CanReadRunTasks:          false,
				CanCreateProject:         false,
			},
			FairRunQueuingEnabled:           true,
			SAMLEnabled:                     false,
			OwnersTeamSAMLRoleID:            nil,
			TwoFactorConformant:             false,
			AssessmentsEnforced:             false,
			DefaultExecutionMode:            "remote",
			RelationshipDefaultAgentPool:    nil,
			RelationshipOAuthTokens:         nil,
			RelationshipAuthenticationToken: nil,
			RelationshipEntitlementSet: &OrganizationListResponseItemRelationshipEntitlementSet{
				ID: "org-iJ5tr4WgB4WpA1hD",
			},
			RelationshipSubscription: nil,
			RelationshipLinksOAuthTokens: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp-two/oauth-tokens",
			},
			RelationshipLinksAuthenticationToken: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp-two/authentication-token",
			},
			RelationshipLinksEntitlementSet: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp-two/entitlement-set",
			},
			RelationshipLinksSubscription: &jsonapi.Links{
				"related": "/api/v2/organizations/hashicorp-two/subscription",
			},
			Links: &jsonapi.Links{
				"self": "/api/v2/organizations/hashicorp-two",
			},
		},
	},
	Links: &jsonapi.Links{
		"self":  "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20",
		"first": "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20",
		"prev":  nil,
		"next":  nil,
		"last":  "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20",
	},
	Meta: &jsonapi.Meta{
		"pagination": map[string]interface{}{
			"current-page": 1,
			"page-size":    20,
			"prev-page":    nil,
			"next-page":    nil,
			"total-pages":  1,
			"total-count":  2,
		},
	},
}

// https://developer.hashicorp.com/terraform/cloud-docs/api-docs/organizations#sample-response
const testOrganizationListResponseMarshalWant = `
{
  "data": [
    {
      "id": "hashicorp",
      "type": "organizations",
      "attributes": {
        "external-id": "org-Hysjx5eUviuKVCJY",
        "created-at": "2021-08-24T23:10:04.675Z",
        "email": "hashicorp@example.com",
        "session-timeout": null,
        "session-remember": null,
        "collaborator-auth-policy": "password",
        "plan-expired": false,
        "plan-expires-at": null,
        "plan-is-trial": false,
        "plan-is-enterprise": false,
        "plan-identifier": "developer",
        "cost-estimation-enabled": true,
        "send-passing-statuses-for-untriggered-speculative-plans": true,
        "allow-force-delete-workspaces": true,
        "name": "hashicorp",
        "permissions": {
          "can-update": true,
          "can-destroy": true,
          "can-access-via-teams": true,
          "can-create-module": true,
          "can-create-team": true,
          "can-create-workspace": true,
          "can-manage-users": true,
          "can-manage-subscription": true,
          "can-manage-sso": true,
          "can-update-oauth": true,
          "can-update-sentinel": true,
          "can-update-ssh-keys": true,
          "can-update-api-token": true,
          "can-traverse": true,
          "can-start-trial": true,
          "can-update-agent-pools": true,
          "can-manage-tags": true,
          "can-manage-varsets": true,
          "can-read-varsets": true,
          "can-manage-public-providers": true,
          "can-create-provider": true,
          "can-manage-public-modules": true,
          "can-manage-custom-providers": false,
          "can-manage-run-tasks": false,
          "can-read-run-tasks": false,
          "can-create-project": true
        },
        "fair-run-queuing-enabled": true,
        "saml-enabled": false,
        "owners-team-saml-role-id": null,
        "two-factor-conformant": false,
        "assessments-enforced": false,
        "default-execution-mode": "remote"
      },
      "relationships": {
        "default-agent-pool": {
            "data": null
        },
        "oauth-tokens": {
          "links": {
            "related": "/api/v2/organizations/hashicorp/oauth-tokens"
          }
        },
        "authentication-token": {
          "links": {
            "related": "/api/v2/organizations/hashicorp/authentication-token"
          }
        },
        "entitlement-set": {
          "data": {
            "id": "org-Hysjx5eUviuKVCJY",
            "type": "entitlement-sets"
          },
          "links": {
            "related": "/api/v2/organizations/hashicorp/entitlement-set"
          }
        },
        "subscription": {
          "links": {
            "related": "/api/v2/organizations/hashicorp/subscription"
          }
        }
      },
      "links": {
        "self": "/api/v2/organizations/hashicorp"
      }
    },
    {
      "id": "hashicorp-two",
      "type": "organizations",
      "attributes": {
        "external-id": "org-iJ5tr4WgB4WpA1hD",
        "created-at": "2022-01-04T18:57:16.036Z",
        "email": "hashicorp@example.com",
        "session-timeout": null,
        "session-remember": null,
        "collaborator-auth-policy": "password",
        "plan-expired": false,
        "plan-expires-at": null,
        "plan-is-trial": false,
        "plan-is-enterprise": false,
        "plan-identifier": "free",
        "cost-estimation-enabled": false,
        "send-passing-statuses-for-untriggered-speculative-plans": false,
        "allow-force-delete-workspaces": false,
        "name": "hashicorp-two",
        "permissions": {
          "can-update": true,
          "can-destroy": true,
          "can-access-via-teams": true,
          "can-create-module": true,
          "can-create-team": false,
          "can-create-workspace": true,
          "can-manage-users": true,
          "can-manage-subscription": true,
          "can-manage-sso": false,
          "can-update-oauth": true,
          "can-update-sentinel": false,
          "can-update-ssh-keys": true,
          "can-update-api-token": true,
          "can-traverse": true,
          "can-start-trial": true,
          "can-update-agent-pools": false,
          "can-manage-tags": true,
          "can-manage-varsets": true,
          "can-read-varsets": true,
          "can-manage-public-providers": true,
          "can-create-provider": true,
          "can-manage-public-modules": true,
          "can-manage-custom-providers": false,
          "can-manage-run-tasks": false,
          "can-read-run-tasks": false,
          "can-create-project": false
        },
        "fair-run-queuing-enabled": true,
        "saml-enabled": false,
        "owners-team-saml-role-id": null,
        "two-factor-conformant": false,
        "assessments-enforced": false,
        "default-execution-mode": "remote"
      },
      "relationships": {
        "default-agent-pool": {
          "data": null
        },
        "oauth-tokens": {
          "links": {
            "related": "/api/v2/organizations/hashicorp-two/oauth-tokens"
          }
        },
        "authentication-token": {
          "links": {
            "related": "/api/v2/organizations/hashicorp-two/authentication-token"
          }
        },
        "entitlement-set": {
          "data": {
            "id": "org-iJ5tr4WgB4WpA1hD",
            "type": "entitlement-sets"
          },
          "links": {
            "related": "/api/v2/organizations/hashicorp-two/entitlement-set"
          }
        },
        "subscription": {
          "links": {
            "related": "/api/v2/organizations/hashicorp-two/subscription"
          }
        }
      },
      "links": {
        "self": "/api/v2/organizations/hashicorp-two"
      }
    }
  ],
  "links": {
    "self": "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20",
    "first": "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20",
    "prev": null,
    "next": null,
    "last": "https://tfe-zone-b0c8608c.ngrok.io/api/v2/organizations?page%5Bnumber%5D=1&page%5Bsize%5D=20"
  },
  "meta": {
    "pagination": {
      "current-page": 1,
      "page-size": 20,
      "prev-page": null,
      "next-page": null,
      "total-pages": 1,
      "total-count": 2
    }
  }
}
`

func TestDefaultOrganizations_List(t *testing.T) {
	type args struct {
		ctx context.Context
		req *OrganizationListRequest
	}
	tests := []struct {
		name    string
		s       *DefaultOrganizations
		args    args
		want    *OrganizationListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultOrganizations{}
			got, err := s.List(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultOrganizations.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultOrganizations.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrganizationListResponse_Marshal(t *testing.T) {
	tests := []struct {
		name    string
		s       *OrganizationListResponse
		want    []byte
		wantErr bool
	}{
		{
			name: "success",
			s:    &testOrganizationListResponseMarshalStruct,
			want: []byte(testOrganizationListResponseMarshalWant),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrganizationListResponse.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.JSONEq(t, string(tt.want), string(got))
		})
	}
}
